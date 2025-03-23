package controllers

import (
	"context"
	"net/http"
	"time"
	"uala-challenge/internal/domain/follows"

	"github.com/gin-gonic/gin"
)

type FollowController struct {
	followRepo follows.FollowRepository
}

func NewFollowController(followRepo follows.FollowRepository) *FollowController {
	return &FollowController{followRepo: followRepo}
}

func (fc *FollowController) CreateFollow(c *gin.Context) {

	var followRequest struct {
		UserId       int `json:"userId" binding:"required"`
		FollowedUser int `json:"followedUser" binding:"required"`
	}

	var err error

	if err = c.ShouldBindJSON(&followRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if followRequest.UserId == followRequest.FollowedUser {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User can't follow itself"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = fc.followRepo.CreateFollow(ctx, followRequest.UserId, followRequest.FollowedUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Follow created successfully"})
}
