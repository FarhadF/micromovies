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
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
    	glog.Error(err)
	}
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		glog.Error("Error EncodingJson in ControllersGetMovies", err)

	}
}

func NewMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	movie := new(models.Movie)
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {
		if movie.Title != "" && movie.Director != "" && movie.Year != "" {
			var resp string
			resp, err = models.NewMovie(movie)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
				w.Write(resp)
				glog.Error(err)
			}else{
				w.WriteHeader(http.StatusOK)
				res := json.RawMessage(`{"id":"` + resp + `"}`)
				w.Write(res)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
			w.Write(resp)
			glog.Error(err)
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	err := models.DeleteMovie(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {
		w.WriteHeader(http.StatusOK)
		res := json.RawMessage(`{"status":"ok"}`)
		w.Write(res)
	}
}