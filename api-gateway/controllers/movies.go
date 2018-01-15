package controllers

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"micromovies/api-gateway/token"
	"net/http"
	"net/http/httputil"
	"net/url"
	"github.com/casbin/casbin"
	"time"
)

//TODO: Find a way to handle http: proxy error: dial tcp 192.168.163.196:8082: getsockopt: connection refused when backend is not available
func ReverseMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	target := &url.URL{Scheme: "http", Host: "192.168.163.196:8081"}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

func ReverseMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	target := &url.URL{Scheme: "http", Host: "192.168.163.196:8081"}
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(w, r)
}

func ReverseMovieProtected(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tokenString, err := token.ExtractToken(r)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {
		parsedToken, tokenStatus := token.ValidateToken(tokenString)
		glog.Info(parsedToken, tokenStatus)
		if tokenStatus == true {
			//false for turn off logging
			// Enable the logging at run-time.
			//e.EnableLog(true)
			//todo: database persistence instead of csv file
			e := casbin.NewEnforcer("./model.conf", "./policy.csv", false)
			sub := parsedToken.Role             // the user that wants to access a resource.
			obj := "/movie/*"         // the resource that is going to be accessed.
			act := " (DELETE)|(POST)" // the operation that the user performs on the resource.
			/*s := time.Now()
			_ = e.Enforce(sub, obj, act)
			glog.Info("casbinbench: ", time.Since(s))*/
			if e.Enforce(sub, obj, act) == true {
				//allow access:
				target := &url.URL{Scheme: "http", Host: "192.168.163.196:8081"}
				proxy := httputil.NewSingleHostReverseProxy(target)
				proxy.ServeHTTP(w, r)
			} else {
				// deny the request, show an error:
				w.Header().Set("Content-Type", "application/json; charset=UTF-8")
				w.WriteHeader(http.StatusForbidden)
				resp := json.RawMessage(`{"status":"forbidden"}`)
				w.Write(resp)
			}
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
