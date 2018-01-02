package main

import (
	"github.com/golang/glog"
	"flag"
	"micromovies/jwt-auth/router"
	"net/http"
	"os"
	"micromovies/jwt-auth/httplog"
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
