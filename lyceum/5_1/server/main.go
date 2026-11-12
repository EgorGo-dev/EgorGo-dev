package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sort"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"lyceum.gitlab.yandexcloud.net/lyceum/go/2026/autumn/course/students/385779-yam-ato-course-1745/5_1/config"
	pb "lyceum.gitlab.yandexcloud.net/lyceum/go/2026/autumn/course/students/385779-yam-ato-course-1745/5_1/pb"
)

type server struct {
	pb.UnimplementedOrderServiceServer
	mu     sync.RWMutex
	orders map[string]*pb.Order
	nextID int
}

func newServer() *server {
	return &server{
		orders: make(map[string]*pb.Order),
		nextID: 1,
	}
}

func (s *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := strconv.Itoa(s.nextID)
	s.nextID++
	s.orders[id] = &pb.Order{Id: id, Item: req.Item, Quantity: req.Quantity}
	return &pb.CreateOrderResponse{Id: id}, nil
}

func (s *server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	order, ok := s.orders[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}
	return &pb.GetOrderResponse{Order: &pb.Order{Id: order.Id, Item: order.Item, Quantity: order.Quantity}}, nil
}

func (s *server) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	order, ok := s.orders[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}
	order.Item = req.Item
	order.Quantity = req.Quantity
	return &pb.UpdateOrderResponse{Order: &pb.Order{Id: order.Id, Item: order.Item, Quantity: order.Quantity}}, nil
}

func (s *server) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.orders[req.Id]; !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}
	delete(s.orders, req.Id)
	return &pb.DeleteOrderResponse{Success: true}, nil
}

func (s *server) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	list := make([]*pb.Order, 0, len(s.orders))
	for _, o := range s.orders {
		list = append(list, &pb.Order{Id: o.Id, Item: o.Item, Quantity: o.Quantity})
	}
	sort.Slice(list, func(i, j int) bool {
		id1, _ := strconv.Atoi(list[i].Id)
		id2, _ := strconv.Atoi(list[j].Id)
		return id1 < id2
	})
	return &pb.ListOrdersResponse{Orders: list}, nil
}

func main() {
	cfg := config.Load()

	grpcLis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Ошибка запуска gRPC: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, newServer())

	go func() {
		log.Printf("gRPC сервер запущен на порту %s", cfg.GRPCPort)
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("Ошибка работы gRPC: %v", err)
		}
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		"localhost:"+cfg.GRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Ошибка подключения gateway к gRPC: %v", err)
	}

	gwMux := runtime.NewServeMux()
	if err := pb.RegisterOrderServiceHandler(context.Background(), gwMux, conn); err != nil {
		log.Fatalf("Ошибка регистрации gateway: %v", err)
	}

	httpServer := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: gwMux,
	}

	go func() {
		log.Printf("HTTP (gateway) сервер запущен на порту %s", cfg.HTTPPort)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка работы HTTP: %v", err)
		}
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	log.Println("Получен сигнал завершения. Начинаем остановку...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("Ошибка остановки HTTP: %v", err)
	}

	grpcServer.GracefulStop()

	fmt.Println("Server Stopped")
}
