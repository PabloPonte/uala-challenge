package tweets

import "context"

// TweetRepository defines the interface for tweet repository operations.
type TweetRepository interface {
	CreateTweet(ctx context.Context, tweet *Tweet) (Tweet, error)
	GetTweetsByUserId(ctx context.Context, userId int) ([]Tweet, error)
}
