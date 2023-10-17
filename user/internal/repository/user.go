package repository

import (
	"context"
	"fmt"

	helperfunc "github.com/muhammedarifp/micro-ecom/user/internal/commonHelp/helperFunc"
	commonhelp "github.com/muhammedarifp/micro-ecom/user/internal/commonHelp/request"
	"github.com/muhammedarifp/micro-ecom/user/internal/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (u *UserRepository) SaveUserIntoDb(ctx context.Context, req *commonhelp.Users) (*domain.Users, error) {
	db := u.DB
	user := domain.Users{
		// CreateAt: time.Now(),
		Name:     req.Name,
		Email:    req.Email,
		Mobile:   req.Mobile,
		Password: helperfunc.PasswordHashing(req.Password),
	}

	tx := db.Create(&user)

	if tx.Error != nil {
		return &domain.Users{}, tx.Error
	}

	return &domain.Users{
		CreateAt: user.CreateAt,
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Mobile:   user.Mobile,
	}, nil
}

func (r *UserRepository) AuthenticateUser(email, pass string) (bool, error) {
	res := r.DB.First(&domain.Users{}, `email = ?`, email)
	if res.Error != nil {
		return false, fmt.Errorf("user not found")
	}

}
