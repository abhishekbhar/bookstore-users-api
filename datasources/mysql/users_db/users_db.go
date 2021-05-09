package users_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Client *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@unix(%s)/%s?charset=utf8",
		"newuser",
		"password",
		"/var/run/mysqld/mysqld.sock",
		"users_db",
	)

	var err error
	Client, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Printf("database sucessfully configured!!")
}