package router

import (
	"uala-challenge/internal/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(tweetController *controllers.TweetController, followController *controllers.FollowController) *gin.Engine {
	router := gin.Default()

	router.POST("/tweet", tweetController.CreateTweet)
	router.POST("/follow", followController.CreateFollow)
	router.GET("/tweet/:userId", tweetController.GetTweets)

	return router
}
