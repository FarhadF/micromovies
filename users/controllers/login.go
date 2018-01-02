package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"micromovies/users/models"
	"encoding/json"
	"github.com/golang/glog"
)

func Login(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	cred := new(models.Credential)
	err := json.NewDecoder(r.Body).Decode(&cred)
	glog.Info(cred)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {
		if cred.Email != "" && cred.Password != "" {
			err := models.Login(cred)
			if err != nil && (err.Error() == "sql: no rows in result set" || err.Error() == "email or password incorrect") {
				w.WriteHeader(http.StatusBadRequest)
				resp := json.RawMessage(`{"error":"email or password incorrect"}`)
				w.Write(resp)
				glog.Error(err)
			} else if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp := json.RawMessage(`{"error":"server failure in auth"}`)
				w.Write(resp)
				glog.Error(err)
			} else {
				//todo: generate jwt token
				resp := json.RawMessage(`{"status":"ok"}`)
				w.WriteHeader(http.StatusOK)
				w.Write(resp)
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			resp := json.RawMessage(`{"error":"fill all the required fields"}`)
			w.Write(resp)
			glog.Error(err)
			}
	}
}
