package controllers

import (
	"encoding/json"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"micromovies/movies/models"
	"net/http"
)

func GetMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	movies, err := models.GetMovies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	}
	if err := json.NewEncoder(w).Encode(movies); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		glog.Error("Error EncodingJson in ControllersGetMovies", err)

	}
}

func GetMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	movie, err := models.GetMovie(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {
		if err := json.NewEncoder(w).Encode(movie); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			glog.Error("Error EncodingJson in ControllersGetMovies", err)

		}
	}
}
//todo: get userid from the request
func NewMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	movie := new(models.Movie)
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
			} else {
				w.WriteHeader(http.StatusOK)
				res := json.RawMessage(`{"id":"` + resp + `"}`)
				w.Write(res)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			resp := json.RawMessage(`{"error":"fill all the required fields"}`)
			w.Write(resp)
			glog.Error(err)
		}
	}
}

func DeleteMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
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
//todo: get userid from the request
func UpdateMovie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	movie := new(models.Movie)
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else if p.ByName("id") != movie.Id {
		w.WriteHeader(http.StatusBadRequest)
		res := json.RawMessage(`{"error":"malformed request"}`)
		w.Write(res)
		glog.Error("id != movie.Id")
	} else {
		err = models.UpdateMovie(movie)
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
}
