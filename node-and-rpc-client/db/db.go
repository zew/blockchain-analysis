package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/zew/gorp"

	_ "github.com/go-sql-driver/mysql"
	// _ "github.com/mattn/go-sqlite3"
)

func InitDb(host, port, dbName, user, pass string, connectionParams map[string]string) *gorp.DbMap {
	// param docu at https://github.com/go-sql-driver/mysql
	paramsJoined := "?"
	for k, v := range connectionParams {
		paramsJoined = fmt.Sprintf("%s%s=%s&", paramsJoined, k, v)
	}
	connStr2 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s%s", user, pass, host, port, dbName, paramsJoined)
	db, err := sql.Open("mysql", connStr2)
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	// dbmap.AddTableWithName(Post{}, "post").SetKeys(true, "Id")
	// err = dbmap.CreateTablesIfNotExists()
	// checkErr(err, "Create tables failed")
	return dbmap
}

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}
