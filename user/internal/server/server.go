package server

import (
	"log"
	"net"

	"github.com/muhammedarifp/micro-ecom/user/internal/proto"
	"github.com/muhammedarifp/micro-ecom/user/internal/repository"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

const (
	PORT = ":9000"
)

func InitializeGrpcServer(db *gorm.DB) {
	lis, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatal(err.Error())
	}

	grpcServer := grpc.NewServer()

	repo := repository.NewRepository(db)
	us := UserServer{
		UserRepo: *repo,
	}

	proto.RegisterUserServiceServer(grpcServer, &us)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err.Error())
	}
}
