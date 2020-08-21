package dbops

import (
	"database/sql"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dbConn *sql.DB
	err error
)

func init()  {
	// dbConn, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/video_server?charset=utf8")
	dbConn, err = sql.Open("sqlite3", "./database/video.db")
	if err != nil {
		panic(err.Error())
	}
}
