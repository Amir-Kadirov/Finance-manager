package main

import (
	"context"
	"finance/api"
	"finance/config"
	"finance/pkg/logger"
	"finance/service"
	"finance/storage/postgres"
	"fmt"
)

func main() {
	cfg := config.Load()
	store, err := postgres.New(context.Background(), cfg)
	if err != nil {
		fmt.Println("error while connecting db, err: ", err)
		return
	}
	defer store.CloseDB()

	service := service.New(store)

	log := logger.New(cfg.ServiceName)

	c := api.New(service, log)

	fmt.Println("programm is running on localhost:8008...")
	c.Run(":8080")
}
