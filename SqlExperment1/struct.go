package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	StuNou   int `db:"id"`
	StuName  string
	StuSex   string
	StuAge   string
	StuGrade string
}

type Course struct {
	CourseID    int
	CourseName  string
	CoursePre   string
	CourseGrade int
}
type SC struct {
	Sno         int
	CourseId    int
	CourseGrade int
}

type SchoolLife interface {
	Insert() error
	Show() error
}

func (stu *Student) Insert() error {

	db, err := InitDb()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	res, err := db.Exec("INSERT INTO STUDENT(SNO,STUNAME,SSEX,STUAGE,STUGRADE) VALUES (?,?,?,?,?)", stu.StuNou, stu.StuName, stu.StuSex, stu.StuAge, stu.StuGrade)

	fmt.Println(res.RowsAffected())

	return err
}

func (course *Course) Insert() error {

	db, err := InitDb()
	if err != nil {
		panic(err.Error())

	}
	defer db.Close()

	res, err := db.Exec("insert into course(COURSEID,COURSENAME,COURSEPRE,COURSEGRADE) VALUES(?,?,?,?) ", course.CourseID, course.CourseName, course.CoursePre, course.CourseGrade)

	fmt.Println(res.RowsAffected())
	return err
}
func (sc *SC) Insert() error {
	db, err := InitDb()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	res, err := db.Exec("insert into sc (SNO,COURSEID,COURSEGRADE) VALUES (?,?,?)", sc.Sno, sc.CourseId, sc.CourseGrade)

	fmt.Println(res.RowsAffected())
	return err

}

func (stu *Student) Show() error {

	db, err := InitDb()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	row, err := db.Query("select * from student ")

	if err != nil {
		panic(err.Error())
	}

	for row.Next() {
		err := row.Scan(&stu.StuNou, &stu.StuName, &stu.StuSex, &stu.StuAge, &stu.StuGrade) //通过这种方式来拿到数据
		if err != nil {
			fmt.Println("查询数据出错")
		}
		fmt.Println(stu)

	}
	return err
}
func (course *Course) Show() error {

	db, err := InitDb()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	row, err := db.Query("select * from course ")

	if err != nil {
		panic(err.Error())
	}

	for row.Next() {
		err := row.Scan(&course.CourseID, &course.CourseName, &course.CoursePre, &course.CourseGrade) //通过这种方式来拿到数据
		if err == nil {
			fmt.Println("查询数据出错")
		}
		fmt.Println(course)

	}
	return err
}
func (sc *SC) Show() error {

	db, err := InitDb()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	row, err := db.Query("select * from student ")

	if err != nil {
		panic(err.Error())
	}

	for row.Next() {
		err := row.Scan(&sc.Sno, &sc.CourseId, &sc.CourseGrade) //通过这种方式来拿到数据
		if err == nil {
			fmt.Println("查询数据出错")
		}
		fmt.Println(sc)

	}
	return err
}

func (stu *Student) UpdateAge() error {

	db, err := InitDb()
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	db.Exec("update student set STUAGE ")


	return nil

}
