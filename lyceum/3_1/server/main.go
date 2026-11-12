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

	pb "lyceum.gitlab.yandexcloud.net/lyceum/go/2026/autumn/course/students/385779-yam-ato-course-1745/3_1/pb"
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

	s.orders[id] = &pb.Order{
		Id:       id,
		Item:     req.Item,
		Quantity: req.Quantity,
	}

	return &pb.CreateOrderResponse{Id: id}, nil
}

func (s *server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	order, ok := s.orders[req.Id]
	if !ok {
		return nil, status.Error(codes.NotFound, "order not found")
	}

	return &pb.GetOrderResponse{Order: &pb.Order{
		Id:       order.Id,
		Item:     order.Item,
		Quantity: order.Quantity,
	}}, nil
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

	return &pb.UpdateOrderResponse{Order: &pb.Order{
		Id:       order.Id,
		Item:     order.Item,
		Quantity: order.Quantity,
	}}, nil
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
		list = append(list, &pb.Order{
			Id:       o.Id,
			Item:     o.Item,
			Quantity: o.Quantity,
		})
	}

	sort.Slice(list, func(i, j int) bool {
		id1, _ := strconv.Atoi(list[i].Id)
		id2, _ := strconv.Atoi(list[j].Id)
		return id1 < id2
	})
	return &pb.ListOrdersResponse{Orders: list}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, newServer())

	log.Println("Server started on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
