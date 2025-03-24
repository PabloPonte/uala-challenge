package tweetsService

import (
	"context"
	"time"
	"uala-challenge/internal/domain/tweets"
	"uala-challenge/internal/services/followService"
)

// interface definition
type TweetsService interface {
	CreateTweet(userId int, content string) (tweets.Tweet, error)
	GetTweetsByUserId(userId int) ([]tweets.Tweet, error)
}

// implementation
type tweetsService struct {
	tweetRepo     tweets.TweetRepository
	followService followService.FollowService
}

func NewTweetsService(tweetRepo tweets.TweetRepository, followService followService.FollowService) TweetsService {
	return &tweetsService{
		tweetRepo:     tweetRepo,
		followService: followService,
	}
}

func (ts *tweetsService) CreateTweet(userId int, content string) (tweets.Tweet, error) {

	tweet := tweets.NewTweet(userId, content)

	if len(tweet.Content) > tweets.MAX_TWEET_LENGTH {
		return tweets.Tweet{}, tweets.ErrTweetTooLong
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return ts.tweetRepo.CreateTweet(ctx, tweet)

}

func (ts *tweetsService) GetTweetsByUserId(userId int) ([]tweets.Tweet, error) {

	// get the users that the user follows

	followers, err := ts.followService.GetFollowersByUserId(userId)

	if err != nil {
		return nil, err
	}

	if len(followers) == 0 {
		return nil, tweets.ErrEmptyTimeline
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tweetsList, err := ts.tweetRepo.GetTweetsByUserId(ctx, userId, followers)

	if err != nil {
		return nil, err
	}

	if len(tweetsList) == 0 {
		return nil, tweets.ErrEmptyTimeline
	}

	return tweetsList, nil
}
