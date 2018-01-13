package main

import (
	"flag"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"micromovies/api-gateway/httplog"
	"micromovies/api-gateway/router"
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
		Addr:    ":7000",
		Handler: loggingHandler,
	}
	glog.Fatal(server.ListenAndServe())

}
