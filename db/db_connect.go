package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var Db *sql.DB

func init() {

	Db, _ = sql.Open("mysql", "root:yzk123...@tcp(127.0.0.1:3306)/my?charset=utf8")
	Db.SetMaxOpenConns(2000)
	Db.SetMaxIdleConns(1000)
	Db.Ping()
}
