package router

import (
	"github.com/julienschmidt/httprouter"
	"imdb/movies/controllers"
)

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/movies", controllers.GetMovies)
	return router
}