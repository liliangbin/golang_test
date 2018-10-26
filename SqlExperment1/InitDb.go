package main

import (
	"database/sql"
	"log"
	"fmt"
)



func InitDb() (*sql.DB, error)   {

	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/liliangbin?charset=utf8")

	err = db.Ping()
	if err != nil {
		log.Println(err)
		fmt.Println("lihai l ")
	}

	return db,err
}