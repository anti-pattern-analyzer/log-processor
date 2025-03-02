package main

import (
	"context"
	"log"
	"log-processor/scheduler"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"log-processor/database"
	"log-processor/kafka"
	"log-processor/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8085"
	}

	database.ConnectDB()
	database.RunMigrations()

	if database.DB == nil {
		log.Fatal("Unable to connect to database")
	}

	go scheduler.StartLogCompletionScheduler()

	reader := kafka.InitializeKafkaReader()
	if reader == nil {
		log.Fatal("Failed to initialize Kafka reader")
	}

	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Println("Starting Kafka Consumer...")
		kafka.ConsumeKafkaMessages(reader)
	}()

	router := gin.Default()
	routes.RegisterRoutes(router)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("\nShutting down server...")

		cancel()
		log.Println("Kafka Consumer Stopped")

		log.Println("Server shutdown complete")
		os.Exit(0)
	}()

	log.Printf("🚀 Server running on http://localhost:%s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
