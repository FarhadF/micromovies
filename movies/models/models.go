package models

import (
	"time"
)

type Movie struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Director  string    `json:"director"`
	Year      string    `json:"year"`
	Userid    string	`json:"userid"`
	CreatedOn time.Time `json:"createdon"`
	UpdatedOn time.Time `json:"updatedon"`
}
