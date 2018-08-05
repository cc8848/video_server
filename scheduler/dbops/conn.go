package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DBConn *sql.DB
	err    error
)

func init() {
	DBConn, err = sql.Open("mysql", "root:jiayou@tcp(localhost:3306)/video_server?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}
