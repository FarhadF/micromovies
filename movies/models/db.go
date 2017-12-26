package models

import (
	"database/sql"
)

func InitDbSession() (db *sql.DB ,err error) {
	db, err = sql.Open("postgres", "postgresql://app_user@192.168.163.196:26257/app_database?sslmode=disable")
	return db,err
}