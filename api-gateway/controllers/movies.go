package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"net/url"
	"net/http/httputil"
)
//TODO: Find a way to handle http: proxy error: dial tcp 192.168.163.196:8082: getsockopt: connection refused when backend is not available
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