package domain

import (
	"time"
)

type Users struct {
	CreateAt time.Time `json:"time"`
	ID       uint64    `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Mobile   string    `json:"mobile"`
	Password string    `json:"password"`
}
