package model

import (
	"log"
	"database/sql"
	"github.com/Zereker/video_server/api/conn"
)

// requests
type User struct {
	Id       int
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInfo struct {
	Id int `json:"id"`
}

func AddUserCredential(username string, password string) error {
	stmt, err := conn.DBConn.Prepare("INSERT INTO user (username, password) VALUES (?, ?)")
	if err != nil {
		log.Printf("AddUserCredential, err: %s", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil {
		log.Printf("AddUserCredential, err: %s", err)
		return err
	}
	return nil
}

func GetUserCredential(username string) (string, error) {
	stmt, err := conn.DBConn.Prepare("select password from user where username = ? and status = 0")
	if err != nil {
		log.Printf("GetUserCredential, err: %s", err)
		return "", err
	}
	defer stmt.Close()
	var password string
	err = stmt.QueryRow(username).Scan(&password)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("GetUserCredential, err: %s", err)
	}
	return password, nil
}

func DeleteUser(username string, password string) error {
	stmt, err := conn.DBConn.Prepare("update user set status = 1 where username = ? and password = ?")
	if err != nil {
		log.Printf("DeleteUser, err: %s", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(username, password)
	if err != nil {
		log.Printf("DeleteUser, err: %s", err)
		return err
	}
	return nil
}

func GetUser(loginName string) (*User, error) {
	stmt, err := conn.DBConn.Prepare("SELECT id, password FROM users WHERE username = ?")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	defer stmt.Close()

	var id int
	var password string

	err = stmt.QueryRow(loginName).Scan(&id, &password)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	res := &User{Id: id, Username: loginName, Password: password}

	return res, nil
}
