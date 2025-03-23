package followsReposiroty

import (
	"context"
	"uala-challenge/internal/domain/follows"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type followRepository struct {
	collection *mongo.Collection
}

func NewFollowRepository(db *mongo.Database) follows.FollowRepository {
	return &followRepository{
		collection: db.Collection("follows"),
	}
}

func (r *followRepository) CreateFollow(ctx context.Context, userId int, followedUser int) (err error) {

	// ## TODO bussines rules, a user can not follow itself

	// using an upsert operation, the follow will be created if it does not exist
	// using $addToSet prevents adding duplicates to the followedUsers array
	filter := bson.M{"userId": userId}
	update := bson.M{"$addToSet": bson.M{"followedUsers": followedUser}}
	opts := options.Update().SetUpsert(true)

	_, err = r.collection.UpdateOne(ctx, filter, update, opts)

	return

}

func (r *followRepository) GetFollwersByUserId(ctx context.Context, userId int) (followers []int, err error) {

	filter := bson.M{"userId": userId}

	result := r.collection.FindOne(ctx, filter)

	// this user is not following anyone
	if result.Err() == mongo.ErrNoDocuments {
		return
	}

	// another error ocurred
	if result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		err = result.Err()
		return
	}

	var follow follows.Follow

	if err = result.Decode(&follow); err != nil {
		return
	}

	followers = follow.FollowedUsers

	return
}
