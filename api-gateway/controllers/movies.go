package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"net/url"
	"net/http/httputil"
)

func ReverseMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	target := &url.URL{Scheme: "http", Host: "192.168.163.196:8081"}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w,r)
}
func ReverseMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	target := &url.URL{Scheme: "http", Host: "192.168.163.196:8081"}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w,r)
}
