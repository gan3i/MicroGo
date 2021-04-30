package app

import (
	"github.com/gan3i/microgo/controllers/ping"
	"github.com/gan3i/microgo/controllers/users"
)

func MapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", users.CreateUser)
	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.SearchUser)
	router.DELETE("/users/:user_id", users.DeleteUser)
	router.PUT("/users", users.UpdateUser)
}
