package controllers

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"micromovies/api-gateway/token"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ReverseUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	target := &url.URL{Scheme: "http", Host: "192.168.163.196:8082"}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}
func ReverseUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	target := &url.URL{Scheme: "http", Host: "192.168.163.196:8082"}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

func ReverseUserProtected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tokenString, err := token.ExtractToken(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {

		_, tokenStatus := token.ValidateToken(tokenString)
		glog.Info(tokenStatus)
		if tokenStatus == true {
			target := &url.URL{Scheme: "http", Host: "192.168.163.196:8082"}
			proxy := httputil.NewSingleHostReverseProxy(target)
			proxy.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusForbidden)
			resp := json.RawMessage(`{"status":"forbidden"}`)
			w.Write(resp)
		}
	}

	/*} else {
		target := &url.URL{Scheme: "http", Host: "192.168.163.196:8082"}
		proxy := httputil.NewSingleHostReverseProxy(target)
		proxy.ServeHTTP(w, r)
	}*/
}
