# API image docker file

###############
# Build Stage #
###############
FROM golang:alpine AS builder
RUN apk --no-cache add build-base
WORKDIR /app

# copy required source files
COPY cmd cmd/
COPY internal internal/
COPY pkg pkg/
COPY go.mod .
COPY go.sum .

# download dependencies
RUN go mod download

# build the app
RUN go build -o tweetAPI cmd/server/main.go

#############
# Run Stage #
#############
FROM alpine:latest
COPY --from=builder /app/tweetAPI tweetAPI
COPY .docker-env .env
ENTRYPOINT ["/tweetAPI"]