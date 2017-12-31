package main

import (
	"github.com/golang/glog"
	"flag"
	_ "github.com/lib/pq"
	"micromovies/api-gateway/router"
	"net/http"
	"os"
	"micromovies/api-gateway/httplog"
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
