package models

import (
	"time"
)

type User struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	LastName string    `json:"lastname"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreatedOn  time.Time `json:"createdon"`
	UpdatedOn  time.Time `json:"updatedon"`
}
/*
CREATE TABLE users (
id UUID NOT NULL DEFAULT gen_random_uuid(),
name STRING NULL,
lastname STRING NULL,
email STRING NOT NULL,
password STRING NOT NULL,
createdon TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
updatedon TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
CONSTRAINT "primary" PRIMARY KEY (id ASC))
*/