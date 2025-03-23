package main

import (
	"context"
	"log"
	"time"
	"uala-challenge/internal/infrastructure/database"
	"uala-challenge/internal/infrastructure/repository"
	"uala-challenge/internal/infrastructure/router"
	"uala-challenge/internal/interfaces/controllers"
	"uala-challenge/pkg/config"
)

const API_VERSION = "0.3.1"

func main() {

	// load the configuration from the .env file
	if err := config.LoadEnvironment(); err != nil {
		log.Fatalf("Could not load environment configuration: %v", err)
	}

	// main context with a 30 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Connect to the MongoDB database and defer the disconnection
	if err := database.Connect(ctx); err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer database.Disconnect(ctx)

	tweetRepo := repository.NewTweetRepository(database.GetDatabase())
	tweetController := controllers.NewTweetController(tweetRepo)

	followRepo := repository.NewFollowRepository(database.GetDatabase())
	followController := controllers.NewFollowController(followRepo)

	// Set up the API routes
	r := router.SetupRouter(tweetController, followController)

	// Start the server
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
