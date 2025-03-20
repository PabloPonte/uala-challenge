package repository

import (
	"context"
	"time"

	"uala-challenge/internal/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type tweetRepository struct {
	collection *mongo.Collection
}

func NewTweetRepository(db *mongo.Database) domain.TweetRepository {

	return &tweetRepository{
		collection: db.Collection("tweets"),
	}
}

func (r *tweetRepository) CreateTweet(ctx context.Context, tweet *domain.Tweet) (domain.Tweet, error) {
	tweet.CreationDate = time.Now()
	result, err := r.collection.InsertOne(ctx, tweet)

	tweet.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return *tweet, err
}

func (r *tweetRepository) GetTweetsByUserId(ctx context.Context, userId int) (tweets []domain.Tweet, err error) {

	// get the users that the user follows
	followers, err := NewFollowRepository(r.collection.Database()).GetFollwersByUserId(ctx, userId)

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
		var tweet domain.Tweet
		if err = cursor.Decode(&tweet); err != nil {
			return
		}
		tweets = append(tweets, tweet)
	}

	return
}
