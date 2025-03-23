package tweets

import "time"

const MAX_TWEET_LENGTH = 280

// Tweet represents a tweet in the system.
type Tweet struct {
	ID           string    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId       int       `json:"userId" bson:"userId"`
	Content      string    `json:"content" bson:"content"`
	CreationDate time.Time `json:"creationDate" bson:"creationDate"`
}

// NewTweet creates a new Tweet instance.
func NewTweet(userId int, content string) *Tweet {
	return &Tweet{
		UserId:  userId,
		Content: content,
	}
}
