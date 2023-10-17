package server

import (
	"context"

	"github.com/muhammedarifp/micro-ecom/cart/internal/proto"
)

type CartServer struct {
	proto.CartServiceServer
	// repo
}

func (*CartServer) AddToCart(ctx context.Context, req *proto.AddCartRequest) (*proto.AddCartResponse, error) {
	return &proto.AddCartResponse{
		Status: false,
	}, nil
}

func (*CartServer) GetAllFromCart(ctx context.Context, req *proto.GetCartRequest) (*proto.GetCartResponse, error) {
	return &proto.GetCartResponse{
		Items: []*proto.Item{
			{Productid: 1, Name: "Ajukka", Price: 111, Quantity: 1},
			{Productid: 2, Name: "Labeebkka", Price: 10000, Quantity: 1},
		},
	}, nil
}
