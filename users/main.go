package main

import (
	"flag"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"micromovies/users/httplog"
	"micromovies/users/models"
	"micromovies/users/router"
	"net/http"
	"os"
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
