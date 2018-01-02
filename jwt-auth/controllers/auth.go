package controllers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/golang/glog"
	"encoding/json"
	"micromovies/jwt-auth/token"
	"micromovies/jwt-auth/models"
)

func CreateToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	tokenReq := new(models.Token)
	err := json.NewDecoder(r.Body).Decode(&tokenReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {
		if tokenReq.Email != "" && tokenReq.Role != "" {
			var resp string
			resp, err = (token.GenerateToken(tokenReq.Email, tokenReq.Role))
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
				w.Write(resp)
				glog.Error(err)
			} else {
				w.WriteHeader(http.StatusOK)
				res := json.RawMessage(`{"token":"` + resp + `"}`)
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