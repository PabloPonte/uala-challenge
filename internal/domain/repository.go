// filepath: /home/pabloponte/Repos/PabloPonte/uala-challenge/internal/domain/repository.go
package domain

import (
	"context"
)

type FollowRepository interface {
	CreateFollow(ctx context.Context, userId int, followedUser int) error
	GetFollwersByUserId(ctx context.Context, userId int) ([]int, error)
}

// TweetRepository defines the interface for tweet repository operations.
type TweetRepository interface {
	CreateTweet(ctx context.Context, tweet *Tweet) (Tweet, error)
	GetTweetsByUserId(ctx context.Context, userId int) ([]Tweet, error)
}
