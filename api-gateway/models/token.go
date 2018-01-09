package models

type AuthToken struct {
	Email	string `json:"email"`
	Exp		int		`json:"exp"`
	Iat 	int		`json:"iat"`
	Role	string	`json:"role"`
}

