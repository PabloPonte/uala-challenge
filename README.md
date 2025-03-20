# Ualá Challenge Tweets API

This project is a simple API for managing tweets and follows using Golang, the Gin framework, and MongoDB. It follows a Domain-Driven Design (DDD) approach to structure the codebase.

## Project Structure

```
uala-challenge
├── cmd
│   └── server
│       └── main.go          # Entry point of the application
├── internal
│   ├── domain
│   │   ├── tweet.go         # Domain model for tweets
│   │   └── follow.go        # Domain model for follows
│   ├── infrastructure
│   │   ├── database
│   │   │   └── mongo.go     # MongoDB connection logic
│   │   ├── router
│   │   │   └── gin.go       # Gin router setup and API endpoints
│   │   └── repository
│   │       ├── tweet_repository.go  # Repository for tweets
│   │       └── follow_repository.go  # Repository for follows
│   ├── interfaces
│   │   ├── controllers
│   │   │   ├── tweet_controller.go   # Controller for tweet-related requests
│   │   │   └── follow_controller.go   # Controller for follow-related requests
│   │   └── repository
│   │       ├── tweet_repository.go    # Interface for tweet repository
│   │       └── follow_repository.go    # Interface for follow repository
│   └── usecases
│       ├── tweet_usecase.go           # Use case for managing tweets
│       └── follow_usecase.go          # Use case for managing follows
├── go.mod                             # Go module file
└── README.md                          # Project documentation
```

## Setup Instructions

1. **Clone the repository:**
   ```
   git clone <repository-url>
   cd uala-challenge
   ```

2. **Install dependencies:**
   Ensure you have Go installed, then run:
   ```
   go mod tidy
   ```

3. **Set up MongoDB:**
   Make sure you have a MongoDB instance running. Update the connection string in `internal/infrastructure/database/mongo.go` if necessary.

4. **Run the application:**
   ```
   go run cmd/server/main.go
   ```

5. **API Usage:**
   - **Create a Tweet:**
     - Endpoint: `POST /tweet/`
     - Body: `{ "user": <userId>, "content": "<tweet content>" }`
   - **Follow a User:**
     - Endpoint: `POST /follow/`
     - Body: `{ "userId": <userId>, "followedUser": <followedUserId> }`
   - **Get User Timeline:**
     - Endpoint: `GET /tweet/:userId`

## Initial Data
pip install pymongo

## run using .env
### building the solution
go build -o tweetapi cmd/server/main.go
./tweetapi
### running using the golang compiler
go run cmd/server/main.go -e $PWD/.env
### debuging in vscode
setting this configuration launch.json
```json
 {
   "name": "Launch Package",
   "type": "go",
   "request": "launch",
   "mode": "auto",
   "program": "cmd/server/main.go",
   "args": ["-e","${workspaceFolder}/.env"],
}
```

## Testing

Unit tests are included in the project. You can run them using:
```
go test ./...
```

## Documentation

Further documentation for the API and data models will be provided in the respective files and can be enhanced using tools like Swagger.

## License

This project is licensed under the MIT License - see the LICENSE file for details.