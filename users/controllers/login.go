package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/golang/glog"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"micromovies/users/models"
	"net/http"
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
			user, err := models.Login(cred)
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
				url := "http://localhost:8083/createtoken"
				var jsonStr = json.RawMessage(`{"email":"` + cred.Email + `","role":"` + user.Role + `"}`)
				//glog.Info(string(jsonStr))
				req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
				req.Header.Set("Content-Type", "application/json")
				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					glog.Error(err)
				}
				defer resp.Body.Close()
				glog.Info("response Status:", resp.Status)
				body, _ := ioutil.ReadAll(resp.Body)
				glog.Info("response Body:", string(body))
				//resp := json.RawMessage(`{"status":"ok"}`)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(body))
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
			resp := json.RawMessage(`{"error":"fill all the required fields"}`)
			w.Write(resp)
			glog.Error(err)
		}
	}
}
