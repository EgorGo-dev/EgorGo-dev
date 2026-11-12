package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"lyceum.gitlab.yandexcloud.net/lyceum/go/2026/autumn/course/students/385779-yam-ato-course-1745/6_1/config"
	pb "lyceum.gitlab.yandexcloud.net/lyceum/go/2026/autumn/course/students/385779-yam-ato-course-1745/6_1/pb"
	"lyceum.gitlab.yandexcloud.net/lyceum/go/2026/autumn/course/students/385779-yam-ato-course-1745/6_1/store"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func main() {
	cfg := config.Load()

	connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := store.New(ctx, connString)
	if err != nil {
		log.Fatalf("DB connection failed: %v", err)
	}
	defer db.Close()

	if err := runMigrations(ctx, db.Pool()); err != nil {
		log.Fatalf("Migrations failed: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, NewServer(db))

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Listen gRPC: %v", err)
	}

	go func() {
		log.Printf("gRPC server on %s", cfg.GRPCPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("gRPC serve: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "localhost:"+cfg.GRPCPort, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Dial gateway: %v", err)
	}

	gwMux := runtime.NewServeMux()
	if err := pb.RegisterOrderServiceHandler(context.Background(), gwMux, conn); err != nil {
		log.Fatalf("Register gateway: %v", err)
	}
	httpServer := &http.Server{Addr: ":" + cfg.HTTPPort, Handler: gwMux}

	go func() {
		log.Printf("HTTP server on %s", cfg.HTTPPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP serve: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down...")
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := httpServer.Shutdown(shutdownCtx); err != nil {
		log.Printf("HTTP shutdown: %v", err)
	}
	grpcServer.GracefulStop()
	fmt.Println("Server Stopped")
}

func runMigrations(ctx context.Context, pool *pgxpool.Pool) error {
	entries, err := fs.ReadDir(migrationsFS, "migrations")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		content, err := fs.ReadFile(migrationsFS, "migrations/"+entry.Name())
		if err != nil {
			return err
		}
		if _, err := pool.Exec(ctx, string(content)); err != nil {
			return fmt.Errorf("migration %s: %w", entry.Name(), err)
		}
	}
	return nil
}
