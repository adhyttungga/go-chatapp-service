package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/adhyttungga/go-chatapp-service/config"
	"github.com/adhyttungga/go-chatapp-service/routes"
)

func main() {
	os.Setenv("TZ", "Asia/Jakarta")
	log.Println("Server initialized...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Mongo DB connection
	connection := config.Connect(ctx)
	defer connection.Disconnect(ctx)

	router := routes.NewRouter(connection.Database("chatapp"))
	// if err := router.Run(":" + config.Config.ServicePort); err != nil {
	// 	log.Fatalf("Listen: %s\n", err)
	// }
	server := &http.Server{
		Addr:    ":" + config.Config.ServicePort,
		Handler: router,
	}

	go func() {
		// Service connections
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("Listen: %s\n", err)
		}
	}()

	log.Printf("Listening and serving HTTP on %s", server.Addr)

	// wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown: %s\n", err)
	}

	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("Timeout of 5 seconds.")
	}

	log.Println("Server Exiting")
}
