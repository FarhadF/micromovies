package router

import (
	"github.com/julienschmidt/httprouter"
	"micromovies/jwt-auth/controllers"
)

func New() *httprouter.Router {
	router := httprouter.New()
	router.POST("/createtoken", controllers.CreateToken)
	//router.POST("/extractclaims", controllers.ExtractClaims)
	return router
}
