package users_db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
)

/*
*
* 		To connect to mysql server running on another container use following commands:
* 		--> Create network 
*				docker network create {network name}
* 		--> Add container to the network
*				docker network connect {network name} {dev container name}
*				docker network connect {network name} {mysql container name}
* 		--> List all the container on the network with ip address
* 				docker network inspect {network name} 
* 
*/



var (
	Client *sql.DB
)

const (
	mysql_users_username  = ""
	mysql_users_password  = ""
	mysql_users_host      = "172.18.0.3" //comes from docker network config
	mysql_users_schema    = ""
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"root",
		"toor",
		mysql_users_host,
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
