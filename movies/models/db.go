package models

import (
	"database/sql"
	"errors"
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

func NewMovie(movie *Movie) (string,error) {
	err := CheckDbSession(db)
	if err != nil {
		return "", err
	}
	rows, err := db.Query("select * from movies where title='" + movie.Title + "'")
	if err != nil {
		return "", err
	}
	if !rows.Next(){
		var  id string
		err := db.QueryRow("insert into movies (title, director, year, userid) values($1,$2,$3,$4) returning id", movie.Title,movie.Director, movie.Year, movie.Userid).Scan(&id)
		//res, err := stmt.Exec(movie.Title,movie.Director, movie.Year, movie.Userid)
		//id, err := res.LastInsertId()
		if err != nil {
			return "", err
		}
		//return strconv.FormatInt(id, 10), nil
		return id, nil
	} else {

		return "", errors.New("movie already exists")
	}
}

func DeleteMovie(id string) error {
	err := CheckDbSession(db)
	if err != nil {
		return err
	}
	rows, err := db.Query("select * from movies where id='" + id + "'")
	if err != nil {
		return err
	}
	if !rows.Next(){
		return errors.New("movie does not exist")
	}
	_ ,err = db.Query("delete from movies where id = $1", id)
		return nil
}