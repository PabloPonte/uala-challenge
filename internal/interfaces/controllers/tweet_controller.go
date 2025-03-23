package controllers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"uala-challenge/internal/domain/tweets"

	"github.com/gin-gonic/gin"
)

type TweetController struct {
	tweetRepo tweets.TweetRepository
}

func NewTweetController(tweetRepo tweets.TweetRepository) *TweetController {
	return &TweetController{tweetRepo: tweetRepo}
}

func (tc *TweetController) CreateTweet(c *gin.Context) {

	var tweetRequest struct {
		UserId  int    `json:"userId" binding:"required"`
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&tweetRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tweet tweets.Tweet

	tweet.UserId = tweetRequest.UserId
	tweet.Content = tweetRequest.Content
	tweet.CreationDate = time.Now()

	if len(tweet.Content) > tweets.MAX_TWEET_LENGTH {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tweet content is too long"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tweet, err := tc.tweetRepo.CreateTweet(ctx, &tweet)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": tweet})
}

func (tc *TweetController) GetTweets(c *gin.Context) {
	userId_param := c.Param("userId")

	userId, err := strconv.Atoi(userId_param)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tweets, err := tc.tweetRepo.GetTweetsByUserId(ctx, userId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(tweets) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No tweets found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"payload": tweets})
}
