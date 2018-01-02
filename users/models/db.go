package models

import (
	"database/sql"
	"errors"
	"time"
	"fmt"
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

func GetUsers() ([]User, error) {
	err := CheckDbSession(db)
	if err != nil {
		return nil , err
	}
	rows,err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	users := make([]User, 0)
	for rows.Next() {
		user := new(User)
		err = rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Password, &user.CreatedOn, &user.UpdatedOn)
		if err != nil {
			return nil, err
		}
		users = append(users, *user)
	}

	return users, nil
}

func GetUser(id string) (User, error) {
	var user User
	err := CheckDbSession(db)
	if err != nil {
		return user , err
	}
	rows := db.QueryRow("select * from users where id = $1", id)
	if err != nil {
		return user, err
	}
	err = rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Password, &user.CreatedOn, &user.UpdatedOn)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetUserByEmail(email string) (User, error) {
	var user User
	err := CheckDbSession(db)
	if err != nil {
		return user , err
	}
	rows := db.QueryRow("select * from users where email = $1", email)
	if err != nil {
		return user, err
	}
	err = rows.Scan(&user.Id, &user.Name, &user.LastName, &user.Email, &user.Password, &user.CreatedOn, &user.UpdatedOn)
	if err != nil {
		return user, err
	}
	return user, nil
}


func NewUser(user *User) (string,error) {
	err := CheckDbSession(db)
	if err != nil {
		return "", err
	}
	rows, err := db.Query("select * from users where email='" + user.Email + "'")
	if err != nil {
		return "", err
	}
	if !rows.Next(){
		var  id string
		err := db.QueryRow("insert into users (name, lastname, email, password) values($1,$2,$3,$4) returning id", user.Name, user.LastName, user.Email, user.Password).Scan(&id)
		if err != nil {
			return "", err
		}
		//return strconv.FormatInt(id, 10), nil
		return id, nil
	} else {

		return "", errors.New("user already exists")
	}
}

func DeleteUser(id string) error {
	err := CheckDbSession(db)
	if err != nil {
		return err
	}
	rows, err := db.Query("select * from users where id='" + id + "'")
	if err != nil {
		return err
	}
	if !rows.Next(){
		return errors.New("user does not exist")
	}
	_ ,err = db.Query("delete from users where id = $1", id)
	if err != nil {
		return err
	}
		return nil
}

func UpdateUser(user *User) error {
	err := CheckDbSession(db)
	if err != nil {
		return err
	}
	rows, err := db.Query("select * from users where id='" + user.Id + "'")
	if err != nil {
		return err
	}
	if !rows.Next(){
		return errors.New("user does not exist")
	}
	updatedon := time.Now()
	fmt.Println(updatedon.Format("2006-01-02 15:04:05.999999"))
	_ ,err = db.Query("update users set name = $1, lastname = $2, email = $3, password = $4 where id = $5, updatedon = $6", user.Name, user.LastName, user.Email, user.Password, updatedon.Format("2017-12-26 05:33:46.689934+00:00"))
	if err != nil {
		return err
	}
	return nil
}

func Login(cred *Credential) error {
	user, err := GetUserByEmail(cred.Email)
	if err != nil {
		return err
	}
	if cred.Password == user.Password {
		return nil
	}
	return errors.New("email or password incorrect")
}