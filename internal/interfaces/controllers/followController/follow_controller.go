package followController

import (
	"net/http"
	"uala-challenge/internal/domain/follows"
	"uala-challenge/internal/services/followService"

	"github.com/gin-gonic/gin"
)

type FollowController struct {
	followService followService.FollowService
}

func NewFollowController(followService followService.FollowService) *FollowController {
	return &FollowController{followService: followService}
}

func (fc *FollowController) CreateFollow(c *gin.Context) {

	var followRequest struct {
		UserId       int `json:"userId" binding:"required"`
		FollowedUser int `json:"followedUser" binding:"required"`
	}

	var err error

	// parameter validation and parsing
	if err = c.ShouldBindJSON(&followRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = fc.followService.FollowUser(followRequest.UserId, followRequest.FollowedUser)

	if err != nil {

		// business error handling
		if err == follows.ErrSelfFollow {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Follow created successfully"})
}
