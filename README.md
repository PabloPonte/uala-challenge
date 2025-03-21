# Ualá Challenge Tweets API v0.2.0

This project is a simple API for managing tweets and user follows.
It's written in Golang using the Gin framework, and the persistence layer is implemented using MongoDB.

## Project Structure

```bash
uala-challenge
├── cmd
│   └── server
│       └── main.go                    # Entry point of the application, server start
├── doc
│   ├── dataabse
│   │   └── collections.md             # Database collections documentation
│   └── API
│       ├── swagger_server.sh          # Swagger UI docker launch script
│       └── swagger.json               # API documentation in Swagger format
├── initial_data
│   ├── initial_data_loader.py         # Script for initial data load
│   └── initial_data.csv               # Inifial data file
├── internal
│   ├── domain
│   │   ├── tweet.go                   # Domain model for tweets
│   │   ├── repository.go              # Repository Interfaces Definitions
│   │   └── follow.go                  # Domain model for follows 
│   ├── infrastructure
│   │   ├── database
│   │   │   └── mongo.go               # MongoDB connection logic
│   │   ├── repository
│   │   │   ├── tweet_repository.go    # Repository implementation for tweets
│   │   │   └── follow_repository.go   # Repository implementation for follows
│   │   └── router
│   │       └── gin.go                 # Gin router setup and API endpoints
│   └── interfaces
│       └── controllers
│           ├── tweet_controller.go    # Controller for tweet-related requests
│           └── follow_controller.go   # Controller for follow-related requests
├── pkg
│   └── config
│       └── config.go                  # Environment configuration handler
├── tests
│   └── test_cases.py                  # Test cases ejecution script
├── .env                               # Environment configuration file 
├── .gitignore                         # Git ignore file 
├── bussiness.txt                      # Business rules and assumptions
├── CHANGELOG.md                       # Changelog File
├── go.mod                             # Go module file
├── go.sum                             # Go module file
├── LICENCE                            # Licence information
└── README.md                          # Project documentation
```

## Prerequisites
In order to run this project you will need the following tools
* **Git** installed, in order to pull this repository.
* **Go** installed, in order to run the project locally, build a binary and/or debug the project.
* **Python** installed, in order to run the test script and the initial data script, some extra python libraries could be needed.
* **Docker** installed, if you want to run the application without instaling any other dependency, you can run the provided docker compose file to run all the services needed.

## Setup Instructions for local run and development

1. **Clone the repository:**
   ```bash
   git clone https://github.com/PabloPonte/uala-challenge.git
   cd uala-challenge
   ```

2. **Install dependencies:**
   Ensure you have Go installed, then run:
   ```bash
   go mod tidy
   ```

3. **Set up MongoDB:**
   Make sure you have a MongoDB instance running. Check that the .env file has the correct variables values.
   
   If you want to run a local instance of MongoDB, you can use the script [Local Database](local_database.sh), to run a docker container.

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

## Build and the application
```bash
# build the application
go build -o tweetapi cmd/server/main.go
# run the application
./tweetapi
```

### Debuging in VSCode
Set this configuration launch.json
```json
 {
   "name": "Launch API Server",
   "type": "go",
   "request": "launch",
   "mode": "auto",
   "program": "cmd/server/main.go",
   "args": ["-e","${workspaceFolder}/.env"],
}
```

## Initial Data
This project includes an initial data file. To load this data into the database, you can run the following:

```bash
cd initial_data
python3 initial_data_loader.py
```

This script assumes that the .env file is in the root folder.
You may need to install some additional libraries to run this script.

## Testing

A python test script is included in the tests folder. This script includes 12 test cases. To run the tests, simply run the script with the application runing:
 in the project. You can run them using:

```bash
cd tests
python3 test_cases.py
```
This script assumes that the application is runing locally in the port 5000.

## Documentation

For database documentation check this file: [Database Info](/doc/database/collections.md)

For the API documentation you check this [Swagger File](/doc/API/swagger.json), you can see the information in a more friendly interface runing a Swagger UI Server:
```bash
cd doc/API
chmod a+x swagger_server.sh
./swagger_server.sh
```
And the going to [localhost:5001](http://localhost:5001) in your browser.


## License

This project is licensed under the MIT License - see the LICENSE file for details.