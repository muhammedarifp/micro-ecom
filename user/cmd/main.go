package main

import (
	"log"

	"github.com/muhammedarifp/micro-ecom/user/internal/config"
	"github.com/muhammedarifp/micro-ecom/user/internal/db"
	"github.com/muhammedarifp/micro-ecom/user/internal/server"
)

const (
	PORT = ":9000"
)

func main() {
	cfg := config.ConfigureConnections()
	db, err := db.ConnectDatabase(cfg)
	if err != nil {
		log.Fatal("db error")
	}

	server.InitializeGrpcServer(db)

}
