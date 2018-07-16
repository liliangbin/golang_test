package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type Person struct {
	UserId   int    `db:"user_id"`
	Username string `db:"username"`
	Sex      string `db:"sex"`
	Email    string `db:"email"`
}

type sql_user struct {
	id       int    `db:"id"`
	username string `db:"username"`
	password string `db:"password"`
	info     string `db:"info"`
}

func main() {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/liliangbin?charset=utf8")
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println(err)
		fmt.Println("lihai l ")
	}

	username := "asdfadf"
	password := "12312241"
	info := "2342342"
	query := []string{"liliasndf", "fasdfas", "23424"}
	//query  = append(query, "23333")  //需要返回给他本身
	for head, value := range query {
		fmt.Println(head, value)

		/*
		在这个的循环中，他会给我们俩个值，一个用于头的分析，一个是具体的值，*/
	}

	res, err := db.Exec("INSERT into sql_user(username,password,info) values (?,?,?)", username, password, info)

	id, err := res.LastInsertId()
	log.Println(id)

	db.QueryRow("select username,password,info from sql_user where id=?", 1).Scan(&username, &password, &info)

	fmt.Println(username, password, info)
	row, err := db.Query("select username from sql_user where id=?", 3)
	for row.Next() {
		err := row.Scan(&username, &password, &info) //通过这种方式来拿到数据
		if err == nil {
			fmt.Println("查询数据出错")
		}


		fmt.Println(username)
	}

	//对应这个地方，我们可以有一个封装。封装成一个函数，最后直接调用就行。
}
