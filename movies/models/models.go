package models

import (
	"time"
)

type Movie struct {
	Id        string    `json:"id"`
	Title     string    `json:"title"`
	Director  string    `json:"director"`
	Year      string    `json:"year"`
	Userid    string    `json:"userid"`
	CreatedOn time.Time `json:"createdon"`
	UpdatedOn time.Time `json:"updatedon"`
}

/*
CREATE TABLE movies (
id UUID NOT NULL DEFAULT gen_random_uuid(),
title STRING NULL,
director STRING NULL,
year STRING NULL,
userid UUID NULL,
createdon TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
updatedon TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
CONSTRAINT "primary" PRIMARY KEY (id ASC),
)*/
