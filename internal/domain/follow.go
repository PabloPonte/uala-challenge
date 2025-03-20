package domain

type Follow struct {
	UserId        int   `json:"userId" bson:"userId"`
	FollowedUsers []int `json:"followedUsers" bson:"followedUsers"`
}
