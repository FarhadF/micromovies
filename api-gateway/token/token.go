package token

import (
	"net/http"
	"strings"
	"errors"
)

func ExtractToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	split := strings.Split(auth, " ")
	if split[0] != "Bearer" || split[1] == "" {

		return "", errors.New("Malformed Auth Header")
	} else {
		return split[1], nil
	}
}
