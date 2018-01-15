package token

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/golang/glog"
	"micromovies/api-gateway/models"
	"net/http"
	"strings"
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

func ValidateToken(tokenStr string) (models.AuthToken, bool) {
	tokenJson := json.RawMessage(`{"token":"` + tokenStr + `"}`)
	url := "http://localhost:8083/validatetoken"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(tokenJson))
	req.Header.Set("Content-Type", "application/json")
	var parsedToken models.AuthToken
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		glog.Error(err)
	}
	defer resp.Body.Close()
	glog.Info("response Status:", resp.Status)
	if resp.Status == "200 OK" {
		err = json.NewDecoder(resp.Body).Decode(&parsedToken)
		glog.Info(parsedToken)
		if err != nil {
			glog.Error(err)
			return parsedToken ,false
		}
		//if parsedToken.Role != requiredClaim {
		//	glog.Info("role is not user")
		//	return parsedToken, false
		//}
		//Expiration is already checked on token.Parse from the jwt-go package
		return parsedToken, true
	}
	return parsedToken, false
}
