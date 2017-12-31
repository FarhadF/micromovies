package router

import (
	"github.com/julienschmidt/httprouter"
	"micromovies/api-gateway/controllers"
)

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/movies", controllers.ReverseMovies)
    router.POST("/movie", controllers.ReverseMovie)
    router.GET("/movie/*get", controllers.ReverseMovie)
    router.DELETE("/movie/*delete", controllers.ReverseMovie)
    router.POST("/movie/*post", controllers.ReverseMovie)
	router.GET("/users", controllers.ReverseUsers)
	router.POST("/user", controllers.ReverseUser)
	router.GET("/user/*get", controllers.ReverseUser)
	router.DELETE("/user/*delete", controllers.ReverseUser)
	router.POST("/user/*post", controllers.ReverseUser)

	return router
}