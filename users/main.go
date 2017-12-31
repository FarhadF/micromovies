package main

import (
	"github.com/golang/glog"
	"flag"
	_ "github.com/lib/pq"
	"micromovies/users/models"
	"micromovies/users/router"
	"net/http"
	"os"
	"micromovies/users/httplog"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	glog.Info("starting...")
	err := models.InitDbSession()
	if err != nil {
		glog.Fatal("error connecting to the database: ", err)
	}
	glog.Info(models.GetUsers())
	router := router.New()
	loggingHandler := logging.NewApacheLoggingHandler(router, os.Stderr)
	server := &http.Server{
		Addr:    ":8082",
		Handler: loggingHandler,
	}
	glog.Fatal(server.ListenAndServe())

}
