package token

import (
	"net/http"
	"strings"
	"errors"
	"encoding/json"
	"github.com/golang/glog"
	"bytes"
	"micromovies/api-gateway/models"
)

func ExtractToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	split := strings.Split(auth, " ")
	if split[0] != "Bearer" || split[1] == "" {

		return "", errors.New("malformed auth header")
	} else {
		return split[1], nil
	}
}

func ValidateToken(tokenStr string, requiredClaim string) bool {
	tokenJson := json.RawMessage(`{"token":"` + tokenStr + `"}`)
	url := "http://localhost:8083/validatetoken"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(tokenJson))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Error(err)
	}
	defer resp.Body.Close()
	glog.Info("response Status:", resp.Status)
	if resp.Status == "200 OK" {
		var parsedToken models.AuthToken
		err = json.NewDecoder(resp.Body).Decode(&parsedToken)
		glog.Info(parsedToken)
		if err != nil {
			glog.Error(err)
			return false
		}
		if parsedToken.Role != requiredClaim {
			glog.Info("role is not user")
			return false
		}
		//Expiration is already checked on token.Parse from the jwt-go package
		return true
	}
	return false
}