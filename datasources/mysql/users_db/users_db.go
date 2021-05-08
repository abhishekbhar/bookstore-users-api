package users_db

import (
	//"os"
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


const (
	mysql_users_username  = "mysql_users_username "
	mysql_users_password  = "mysql_users_password"
	mysql_users_host      = "mysql_users_host" 		//comes from docker network config
	mysql_users_schema    = "mysql_users_schema"
)


var (
	Client *sql.DB
	// username = os.Getenv(mysql_users_username)
	// password = os.Getenv(mysql_users_password)
	// host 	 = os.Getenv(mysql_users_host)
	// schema 	 = os.Getenv(mysql_users_schema)
	username = "root"
	password = "toor"
	host 	 = "127.18.0.3"
	schema 	 = "users_db"
	
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",username,password,host,schema)

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
