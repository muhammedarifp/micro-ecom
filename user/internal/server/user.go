package server

import (
	"context"
	"fmt"

	"github.com/go-playground/validator/v10"
	commonhelp "github.com/muhammedarifp/micro-ecom/user/internal/commonHelp/request"
	"github.com/muhammedarifp/micro-ecom/user/internal/proto"
	"github.com/muhammedarifp/micro-ecom/user/internal/repository"
)

type UserServer struct {
	proto.UserServiceServer
	UserRepo repository.UserRepository
}

func (s *UserServer) SignupUser(ctx context.Context, req *proto.SignupRequest) (*proto.SignupResponse, error) {
	userData := commonhelp.Users{
		Name:     req.Name,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: req.Password,
	}

	validator := validator.New()
	if validErr := validator.Struct(&userData); validErr != nil {
		return &proto.SignupResponse{
			Status: false,
			Error:  "entered data is incorrect",
			User:   &proto.UserResponse{},
		}, fmt.Errorf("validation error | ", validErr.Error())
	}

	user, err := s.UserRepo.SaveUserIntoDb(context.Background(), &userData)

	if err != nil {
		return &proto.SignupResponse{
			Status: false,
			Error:  "internal server error",
			User:   &proto.UserResponse{},
		}, err
	}

	return &proto.SignupResponse{
		Status: true,
		Error:  "",
		User: &proto.UserResponse{
			Created: user.CreateAt.GoString(),
			Userid:  "1",
			Name:    user.Name,
			Email:   user.Email,
		},
	}, nil

}

// Login function

func (*UserServer) LoginUser(ctx context.Context, req *proto.LoginRequest) (*proto.UserResponse, error) {
	return &proto.UserResponse{
		Userid:  "1",
		Name:    "sssss",
		Email:   "q",
		Created: "sss",
	}, nil
}
