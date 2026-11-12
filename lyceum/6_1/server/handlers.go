package main

import (
	"context"
	"strconv"

	"github.com/jackc/pgx/v5"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "lyceum.gitlab.yandexcloud.net/lyceum/go/2026/autumn/course/students/385779-yam-ato-course-1745/6_1/pb"
	"lyceum.gitlab.yandexcloud.net/lyceum/go/2026/autumn/course/students/385779-yam-ato-course-1745/6_1/store"
)

type Server struct {
	pb.UnimplementedOrderServiceServer
	store *store.Store
}

func NewServer(s *store.Store) *Server {
	return &Server{store: s}
}

func (s *Server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	id, err := s.store.CreateOrder(ctx, req.Item, int(req.Quantity))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create: %v", err)
	}
	return &pb.CreateOrderResponse{Id: strconv.Itoa(id)}, nil
}

func (s *Server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "order not found")
	}
	item, quantity, err := s.store.GetOrder(ctx, id)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, status.Error(codes.NotFound, "order not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to get: %v", err)
	}
	return &pb.GetOrderResponse{Order: &pb.Order{Id: req.Id, Item: item, Quantity: int32(quantity)}}, nil
}

func (s *Server) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "order not found")
	}
	err = s.store.UpdateOrder(ctx, id, req.Item, int(req.Quantity))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update: %v", err)
	}
	return &pb.UpdateOrderResponse{Order: &pb.Order{Id: req.Id, Item: req.Item, Quantity: req.Quantity}}, nil
}

func (s *Server) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
	id, err := strconv.Atoi(req.Id)
	if err != nil {
		return nil, status.Error(codes.NotFound, "order not found")
	}
	err = s.store.DeleteOrder(ctx, id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete: %v", err)
	}
	return &pb.DeleteOrderResponse{Success: true}, nil
}

func (s *Server) ListOrders(ctx context.Context, req *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	orders, err := s.store.ListOrders(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list: %v", err)
	}
	var pbOrders []*pb.Order
	for _, o := range orders {
		pbOrders = append(pbOrders, &pb.Order{Id: strconv.Itoa(o.ID), Item: o.Item, Quantity: int32(o.Quantity)})
	}
	return &pb.ListOrdersResponse{Orders: pbOrders}, nil
}
