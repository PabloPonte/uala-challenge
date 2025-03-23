package follows

import "context"

type FollowRepository interface {
	CreateFollow(ctx context.Context, userId int, followedUser int) error
	GetFollwersByUserId(ctx context.Context, userId int) ([]int, error)
}
