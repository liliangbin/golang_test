package main

import (
	"testing"
	"fmt"
)

func TestCourse_Insert(t *testing.T) {

	db, err := InitDb()
	if err != nil {

		panic(err.Error())
	}
	defer db.Close()


	fmt.Println()
	res, err := db.Exec("insert into student (SNO,STUNAME,SSEX,STUAGE,STUGRADE) VALUES(?,?,?,?,?)", )
}
