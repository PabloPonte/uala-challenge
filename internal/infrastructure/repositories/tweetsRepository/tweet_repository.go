package tweetsRepository

import (
	"context"
	"time"
	"uala-challenge/internal/domain/tweets"
	"uala-challenge/internal/infrastructure/repositories/followsReposiroty"
	"uala-challenge/internal/services/followService"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type tweetRepository struct {
	collection *mongo.Collection
}

func NewTweetRepository(db *mongo.Database) tweets.TweetRepository {

	return &tweetRepository{
		collection: db.Collection("tweets"),
	}
}

func (r *tweetRepository) CreateTweet(ctx context.Context, tweet *tweets.Tweet) (tweets.Tweet, error) {
	tweet.CreationDate = time.Now()
	result, err := r.collection.InsertOne(ctx, tweet)

	tweet.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return *tweet, err
}

func (r *tweetRepository) GetTweetsByUserId(ctx context.Context, userId int) (tweetsList []tweets.Tweet, err error) {

	// get the users that the user follows

	followService := followService.NewFollowService(followsReposiroty.NewFollowRepository(r.collection.Database()))

	followers, err := followService.GetFollowersByUserId(userId)

	if err != nil {
		return
	}

	if len(followers) == 0 {
		return
	}

	// get the tweets of the users that the user follows
	filter := bson.M{"userId": bson.M{"$in": followers}}

	// set the results order by creation date
	opts := options.Find().SetSort(bson.D{{Key: "creationDate", Value: -1}})

	cursor, err := r.collection.Find(ctx, filter, opts)

	if err != nil {
		return
	}

	defer cursor.Close(ctx)

	// get all the resutls
	for cursor.Next(ctx) {
		var tweet tweets.Tweet
		if err = cursor.Decode(&tweet); err != nil {
			return
		}
		tweetsList = append(tweetsList, tweet)
	}

	return
}
