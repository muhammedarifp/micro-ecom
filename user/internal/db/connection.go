package db

import (
	"fmt"

	"github.com/muhammedarifp/micro-ecom/user/internal/config"
	"github.com/muhammedarifp/micro-ecom/user/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cnfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s", cnfg.DBUser, cnfg.DBPassword, cnfg.DBName, cnfg.DBPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if miErr := db.AutoMigrate(
		&domain.Users{},
	); miErr != nil {
		return nil, fmt.Errorf(miErr.Error())
	}

	return db, nil
}
