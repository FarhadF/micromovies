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
	router.DELETE("/movie/*delete", controllers.ReverseMovieProtected)
	router.POST("/movie/*post", controllers.ReverseMovieProtected)
	router.GET("/users", controllers.ReverseUsers)
	router.POST("/user", controllers.ReverseUser)
	router.GET("/user/id/*get", controllers.ReverseUserProtected)
	router.GET("/user/email/*get", controllers.ReverseUserProtected)
	router.DELETE("/user/*delete", controllers.ReverseUserProtected)
	router.POST("/user/*post", controllers.ReverseUser)
	router.POST("/login", controllers.ReverseUser)

	return router
}
