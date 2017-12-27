package router

import (
	"github.com/julienschmidt/httprouter"
	"micromovies/movies/controllers"
)

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/movies", controllers.GetMovies)
	router.POST("/movie", controllers.NewMovie)
	router.DELETE("/movie/:id", controllers.DeleteMovie)
	router.POST("/movie/:id", controllers.UpdateMovie)
	return router
}