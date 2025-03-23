package router

import (
	"uala-challenge/internal/interfaces/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(tweetController *controllers.TweetController, followController *controllers.FollowController) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	// Enable CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all origins
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.POST("/tweet", tweetController.CreateTweet)
	router.POST("/follow", followController.CreateFollow)
	router.GET("/tweet/:userId", tweetController.GetTweets)

	return router
}
