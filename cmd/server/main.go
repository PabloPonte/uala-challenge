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

const API_VERSION = "0.3.0"

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

	tweet_repo := repository.NewTweetRepository(database.GetDatabase())
	tweet_controller := controllers.NewTweetController(tweet_repo)

	follow_repo := repository.NewFollowRepository(database.GetDatabase())
	follow_controller := controllers.NewFollowController(follow_repo)

	// Set up the API routes
	r := router.SetupRouter(tweet_controller, follow_controller)

	// Start the server
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
