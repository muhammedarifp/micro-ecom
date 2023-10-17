package server

import (
	"log"
	"net"

	pb "github.com/muhammedarifp/micro-ecom/cart/internal/proto"
	"google.golang.org/grpc"
)

func ConnectGrpcServer() {
	lis, err := net.Listen("tcp", ":9001")
	if err != nil {
		log.Fatal("net.Listen error found | ", err.Error())
	}

	// Create grpc server
	grpcServer := grpc.NewServer()

	pb.RegisterCartServiceServer(grpcServer, &CartServer{})

	grpcServer.Serve(lis)
}
