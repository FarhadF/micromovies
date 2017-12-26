package models

import (
	"database/sql"
)
var db *sql.DB
func InitDbSession() (error) {
	var err error
	db, err = sql.Open("postgres", "postgresql://app_user@192.168.163.196:26257/app_database?sslmode=disable")
	return err
}

func CheckDbSession(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		err = InitDbSession()
		if err != nil {
			return err
		} else {
			return nil
		}
	} else {
		return nil
	}
}

func GetMovies() ([]Movie, error) {
	err := CheckDbSession(db)
	if err != nil {
		return nil , err
	}
	rows,err := db.Query("select * from movies")
	if err != nil {
		return nil, err
	}
	movies := make([]Movie, 0)
	for rows.Next() {
		movie := new(Movie)
		err = rows.Scan(&movie.Id, &movie.Title, &movie.Director, &movie.Year, &movie.Userid, &movie.CreatedOn, &movie.UpdatedOn)
		if err != nil {
			return nil, err
		}
		movies = append(movies, *movie)
	}

	return movies, nil
}