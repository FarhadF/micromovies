package main

import (
	"flag"
	"github.com/golang/glog"
	"micromovies/jwt-auth/httplog"
	"micromovies/jwt-auth/router"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	glog.Info("starting...")
	router := router.New()
	loggingHandler := logging.NewApacheLoggingHandler(router, os.Stderr)
	server := &http.Server{
		Addr:    ":8083",
		Handler: loggingHandler,
	}
	glog.Fatal(server.ListenAndServe())
}
