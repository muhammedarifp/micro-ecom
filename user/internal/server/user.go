package server

import (
	"context"
	"log"
	"strconv"

	helperfunc "github.com/muhammedarifp/micro-ecom/user/internal/commonHelp/helperFunc"
	commonhelp "github.com/muhammedarifp/micro-ecom/user/internal/commonHelp/request"
	"github.com/muhammedarifp/micro-ecom/user/internal/proto"
	"github.com/muhammedarifp/micro-ecom/user/internal/repository"
)

type UserServer struct {
	proto.UserServiceServer
	UserRepo repository.UserRepository
}

func (s *UserServer) SignupUser(ctx context.Context, req *proto.SignupRequest) (*proto.UserResponse, error) {
	if len(req.Name) <= 1 {
		log.Println("nil error found")
		return &proto.UserResponse{}, nil
	}

	user, err := s.UserRepo.SaveUserIntoDb(context.Background(), &commonhelp.Users{
		Name:     req.Name,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: helperfunc.PasswordHashing(req.Password),
	})

	if err != nil {
		return &proto.UserResponse{}, err
	} else {
		return &proto.UserResponse{
			Created: user.CreateAt.String(),
			Name:    user.Name,
			Userid:  strconv.Itoa(int(user.ID)),
			Email:   user.Email,
		}, err
	}
}
