package main

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func main() {

	stu := Student{95006, "李良彬", "男", "18", "2018-2019"}

	stu.Show()

	//stu.Insert()

	fmt.Println(stu)



}
