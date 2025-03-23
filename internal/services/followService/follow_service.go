package followService

import (
	"context"
	"time"
	"uala-challenge/internal/domain/follows"
)

// interface definition
type FollowService interface {
	FollowUser(userId int, followedUserId int) error
	GetFollowersByUserId(userId int) ([]int, error)
}

// implementation
type followService struct {
	followRepo follows.FollowRepository
}

func NewFollowService(followRepo follows.FollowRepository) FollowService {
	return &followService{followRepo: followRepo}
}

func (fs *followService) FollowUser(userId int, followedUser int) error {

	if userId == followedUser {
		return follows.ErrSelfFollow
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return fs.followRepo.CreateFollow(ctx, userId, followedUser)
}

func (fs *followService) GetFollowersByUserId(userId int) ([]int, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	return fs.followRepo.GetFollwersByUserId(ctx, userId)
}
