package models

type Token struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

type TokenRec struct {
	TokenString string `json:"token"`
}
