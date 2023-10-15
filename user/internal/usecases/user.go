package usecases

// import (
// 	"context"

// 	commonhelp "github.com/muhammedarifp/micro-ecom/user/internal/commonHelp/request"
// 	"github.com/muhammedarifp/micro-ecom/user/internal/repository"
// )

// type UserUseCase struct {
// 	UserRepository repository.UserRepository
// }

// func NewUserUsecase(repo repository.UserRepository) *UserUseCase {
// 	return &UserUseCase{
// 		UserRepository: repo,
// 	}
// }

// func (uc *UserUseCase) SignupUserUsecase(ctx context.Context, req commonhelp.Users) error {
// 	return uc.UserRepository.SignupUserRepo(context.Background(), req)
// }
