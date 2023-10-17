package domain

import (
	"time"
)

type Users struct {
	CreateAt time.Time `json:"time,omitempty" validate="required"`
	ID       uint64    `json:"id" gorm:"primaryKey"`
	Name     string    `json:"name" validate="required"`
	Email    string    `json:"email" validate="email"`
	Mobile   string    `json:"mobile" validate="mobile"`
	Password string    `json:"password" validate="required min=5"`
}
