package tweetController

import (
	"net/http"
	"strconv"
	"uala-challenge/internal/domain/tweets"
	"uala-challenge/internal/services/tweetsService"

	"github.com/gin-gonic/gin"
)

type TweetController struct {
	tweetService tweetsService.TweetsService
}

func NewTweetController(tweetService tweetsService.TweetsService) *TweetController {
	return &TweetController{tweetService: tweetService}
}

func (tc *TweetController) CreateTweet(c *gin.Context) {

	var tweetRequest struct {
		UserId  int    `json:"userId" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	// parameter validation and parsing
	if err := c.ShouldBindJSON(&tweetRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweet, err := tc.tweetService.CreateTweet(tweetRequest.UserId, tweetRequest.Content)

	if err != nil {

		// business error handling
		if err == tweets.ErrTweetTooLong {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": tweet})
}

func (tc *TweetController) GetTweets(c *gin.Context) {

	userId_param := c.Param("userId")

	// parameter parse and validation
	userId, err := strconv.Atoi(userId_param)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tweetList, err := tc.tweetService.GetTweetsByUserId(userId)

	if err != nil {

		// business error handling
		if err == tweets.ErrEmptyTimeline {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payload": tweetList})
}
