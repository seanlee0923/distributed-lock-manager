package mysql

import (
	"database/sql"
	"dlm/config"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func Connect(config config.Config) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.MysqlUser, config.MysqlPassword, config.MysqlHost, config.MysqlPort, config.MysqlDatabase)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	Db = db

}
