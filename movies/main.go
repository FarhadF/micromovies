package main

import (
	"flag"
	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"micromovies/movies/httplog"
	"micromovies/movies/models"
	"micromovies/movies/router"
	"net/http"
	"os"
)

func main() {
	flag.Parse()
	defer glog.Flush()
	glog.Info("starting...")
	//db, err := sql.Open("postgres", "postgresql://app_user@192.168.163.196:26257/app_database?sslmode=disable")
	err := models.InitDbSession()
	if err != nil {
		glog.Fatal("error connecting to the database: ", err)
	}
	/*glog.Info(db)
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS movies (id UUID PRIMARY KEY DEFAULT gen_random_uuid() , title VARCHAR , director VARCHAR, year VARCHAR, userid UUID, createdon timestamp with time zone DEFAULT now() NOT NULL, updatedon timestamp with time zone DEFAULT now() NOT NULL)"); err != nil {
		glog.Fatal(err)
	}
	glog.Info("db ping: " ,db.Ping())
	glog.Info(models.CheckDbSession(db))
	rows, err := db.Query("select * from movies")
	defer rows.Close()
	if err != nil {
		glog.Fatal("err", err)
	}
	movie := new(models.Movie)
	for rows.Next() {
		glog.Info("there was a row")
		if err := rows.Scan(&movie.Id, &movie.Title, &movie.Director, &movie.Year, &movie.Userid, &movie.CreatedOn, &movie.UpdatedOn); err != nil {
			glog.Fatal(err)
		}
		glog.Info("movie: ", movie)
	}*/
	glog.Info(models.GetMovies())
	router := router.New()
	loggingHandler := logging.NewApacheLoggingHandler(router, os.Stderr)
	server := &http.Server{
		Addr:    ":8081",
		Handler: loggingHandler,
	}
	glog.Fatal(server.ListenAndServe())

}
