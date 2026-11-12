package main

import (
	"context"
	"log"
	"net"
	"sort"
	"strconv"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"lyceum.gitlab.yandexcloud.net/lyceum/go/2026/autumn/course/students/385779-yam-ato-course-1745/4_1/config"
	pb "lyceum.gitlab.yandexcloud.net/lyceum/go/2026/autumn/course/students/385779-yam-ato-course-1745/4_1/pb"
)

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Printf("[gRPC] Получен запрос: %s", info.FullMethod)

	resp, err := handler(ctx, req)

	if err != nil {
		log.Printf("[gRPC] Ошибка в %s: %v", info.FullMethod, err)
	}
	return resp, err
}

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

	lis, err := net.Listen("tcp", ":"+cfg.GRPCPort)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)
	pb.RegisterOrderServiceServer(grpcServer, newServer())

	log.Printf("Сервер запущен на порту %s", cfg.GRPCPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Ошибка работы сервера: %v", err)
	}
}
