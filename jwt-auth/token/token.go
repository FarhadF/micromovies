package token

import (
	//"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	//"log"
	"net/http"
	"strings"
	"time"
)

type errOut struct {
	Error string `json:error`
}

const mySigningKey = "Super_Dup3r_S3cret"

func GenerateToken(email string, role string) (string, error) {
	// Create the token
	tokenObject := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	tokenObject.Claims = jwt.MapClaims{
		"exp":   time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		"iat":   time.Now().Unix(),
		"email": email,
		"role":  role,
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := tokenObject.SignedString([]byte(mySigningKey))
	return tokenString, err
}

func ParseToken(myToken string) (*jwt.Token, error) {
	parsedToken, err := jwt.Parse(myToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(mySigningKey), nil
	})
	//fmt.Println(token.Claims)
	if err == nil && parsedToken.Valid {
		return parsedToken, nil
	} else {
		return &jwt.Token{}, err
	}
}

func ExtractToken(r *http.Request) (string, error) {
	auth := r.Header.Get("Authorization")
	split := strings.Split(auth, " ")
	if split[0] != "Bearer" || split[1] == "" {

		return "", errors.New("Malformed Auth Header")
	} else {
		return split[1], nil
	}

}

/*func TokenHandler(w http.ResponseWriter, r *http.Request) bool {

	authToken, err := ExtractToken(r)
	if err != nil {
		errout := new(errOut)
		errout.Error = err.Error()
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(errout); err != nil {
			log.Panic("Error EncodingJson in TokenHandler", err)
		}
		return false
	} else {
		parsedToken, err := ParseToken(authToken)
		//fmt.Println("tokenStatus: ", tokenStatus, "err: ", err.Error())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Panic("Error EncodingJson in TokenHandler", err)
			log.Println("token status err: ", err)
			return false

		} else {

			//w.Header().Set("Access-Control-Allow-Origin", "*")
			return true

		}
	}

}
*/
