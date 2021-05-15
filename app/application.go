package app


import (
	"github.com/gin-gonic/gin"
	"github.com/abhishekbhar/bookstore-users-api/logger"
)


var (
	router = gin.Default()
)


func StartApp(){
	mapUrls()
	logger.Info("about to start the application....")
	router.Run(":8081")
}