package follows

import "errors"

// domain errors
var ErrSelfFollow = errors.New("User can't follow itself")

type Follow struct {
	UserId        int   `json:"userId" bson:"userId"`
	FollowedUsers []int `json:"followedUsers" bson:"followedUsers"`
}
