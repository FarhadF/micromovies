package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/golang/glog"
	"encoding/json"
	"micromovies/users/models"
)

func GetUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
    users,err := models.GetUsers()
    if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
    	glog.Error(err)
	}
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		glog.Error("Error EncodingJson in ControllersGetMovies", err)

	}
}

func GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	user,err := models.GetUser(p.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {
		if err := json.NewEncoder(w).Encode(user); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			glog.Error("Error EncodingJson in ControllersGetMovies", err)

		}
	}
}

func NewUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	user := new(models.User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {
		if user.Name != "" && user.LastName != "" && user.Email != "" && user.Password != "" {
			var resp string
			resp, err = models.NewUser(user)
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
			resp := json.RawMessage(`{"error":"fill all the required fields"}`)
			w.Write(resp)
			glog.Error(err)
		}
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	err := models.DeleteUser(p.ByName("id"))
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

func UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	user := new(models.User)
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else if p.ByName("id") != user.Id {
		w.WriteHeader(http.StatusBadRequest)
		res := json.RawMessage(`{"error":"malformed request"}`)
		w.Write(res)
		glog.Error("id != user.Id")
	} else {
		err = models.UpdateUser(user)
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