package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"net/url"
	"net/http/httputil"
	"micromovies/api-gateway/token"
	"github.com/golang/glog"
	"encoding/json"
)

func ReverseUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	target := &url.URL{Scheme: "http", Host: "192.168.163.196:8082"}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w,r)
}
func ReverseUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	target := &url.URL{Scheme: "http", Host: "192.168.163.196:8082"}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w,r)
}

func ReverseUserProtected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tokenString, err := token.ExtractToken(r)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {

		target := &url.URL{Scheme: "http", Host: "192.168.163.196:8082"}
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.ServeHTTP(w, r)
	}
}