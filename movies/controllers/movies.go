package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/golang/glog"
	"encoding/json"
	"imdb/movies/models"
)

func GetMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
    movies,err := models.GetMovies()
    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp, _ := json.Marshal(`{"error":"` + `err.Error()"}`)
		w.Write(resp)
    	glog.Error(err)
	}
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		glog.Error("Error EncodingJson in ControllersGetMovies", err)

	}
}
