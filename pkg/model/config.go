package model

import (
	"fmt"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var databaseName, databaseURL string

func InitDatabase(dbName, dbURL string)  {
	databaseName = dbName
	databaseURL = dbURL
}

func getConnection() *sql.DB {
	dbConnect, err := sql.Open(databaseName, databaseURL)
	if err != nil {
		log.Fatalf("can not connect database : %v", err)
	}
	return dbConnect
}

func CreateDatabase() {
	sqlCreateTable := `
	CREATE TABLE APP_USER (
	ID INT PRIMARY KEY AUTO_INCREMENT,
	employee_id varchar(255) NOT NULL UNIQUE,
	username varchar(255) DEFAULT NULL,
	password varchar(255) DEFAULT NULL,
	name varchar(255) DEFAULT NULL,
	sex varchar(255) DEFAULT NULL,
	department varchar(255) DEFAULT NULL,
	email varchar(255) DEFAULT NULL,
	telephone varchar(255) DEFAULT NULL,
	status varchar(255) DEFAULT NULL,
	createDate datetime DEFAULT NULL,
	updateDate datetime DEFAULT NULL);`

	db := getConnection()
	defer db.Close()
	_, err := db.Exec(sqlCreateTable)
	checkErr(err)
	fmt.Println("Create Table...")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}