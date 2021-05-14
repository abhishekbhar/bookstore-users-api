package app

import (
	"github.com/abhishekbhar/bookstore-users-api/controllers/ping"
	"github.com/abhishekbhar/bookstore-users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	
	router.GET("internal//users/search",      users.Search)
	router.POST("/users", 		     users.Create)	
	router.GET("/users/:user_id",    users.Get	 )
	router.PUT("/users/:user_id",    users.Update)
	router.PATCH("/users/:user_id",  users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.POST("/users/login",		 users.Login)
}