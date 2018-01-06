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
			resp, err = token.GenerateToken(tokenReq.Email, tokenReq.Role)
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

func ValidateToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	tokenRec := new(models.TokenRec)
	err := json.NewDecoder(r.Body).Decode(&tokenRec)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
		w.Write(resp)
		glog.Error(err)
	} else {
		if tokenRec.TokenString != "" {
			parsedToken, err := token.ParseToken(tokenRec.TokenString)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp := json.RawMessage(`{"error":"` + err.Error() + `"}`)
				w.Write(resp)
				glog.Error(err)
			} else if parsedToken.Valid == false {
				w.WriteHeader(http.StatusBadRequest)
				resp := json.RawMessage(`{"error":"invalid token"}`)
				w.Write(resp)
			} else {
				w.WriteHeader(http.StatusOK)
				err := json.NewEncoder(w).Encode(parsedToken.Claims)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					glog.Error(err)
				}
				glog.Info(parsedToken.Claims)
			}

		}
	}

}