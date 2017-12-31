package router

import (
	"github.com/julienschmidt/httprouter"
	"micromovies/users/controllers"
)

func New() *httprouter.Router {
	router := httprouter.New()
	router.GET("/users", controllers.GetUsers)
	router.POST("/user", controllers.NewUser)
	router.DELETE("/user/:id", controllers.DeleteUser)
	router.POST("/user/:id", controllers.UpdateUser)
	router.GET("/user/:id", controllers.GetUser)
	return router
}
