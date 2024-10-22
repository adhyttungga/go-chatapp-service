package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/adhyttungga/go-chatapp-service/config"
	"github.com/adhyttungga/go-chatapp-service/routes"
	"github.com/go-playground/validator/v10"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	connection := config.Connect(ctx)
	validate := validator.New()
	defer connection.Disconnect(ctx)
	defer cancel()

	router := routes.NewRouter(connection.Database("chatapp"), validate)
	if err := router.Run(":" + config.Config.ServicePort); err != nil {
		log.Fatalf("Listen: %s\n", err)
	}
}
