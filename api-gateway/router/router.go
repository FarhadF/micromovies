package router

import (
	"github.com/julienschmidt/httprouter"
	"micromovies/api-gateway/controllers"
)

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/movies", controllers.ReverseMovies)
    router.POST("/movie", controllers.ReverseMovie)
    router.GET("/movie/*wtf", controllers.ReverseMovie)
    router.DELETE("/movie/*delete", controllers.ReverseMovie)
    router.POST("/movie/*post", controllers.ReverseMovie)

//	router.DELETE("/movie/:id", controllers.DeleteMovie)
//	router.POST("/movie/:id", controllers.UpdateMovie)
	return router
}