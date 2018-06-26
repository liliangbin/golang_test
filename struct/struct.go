package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

/*
在这个时候，我们能够体会到，一个struct可以被做成一个小型的类*/
type Student struct {
	Human
	speciality string
}

/*
我们同样是可以通过传输一个struct来出传递一个数据对象的*/
func fallback(p1 Student, p2 Student) (Student, int) {

	return p1, p2.age
}

type Rectangle struct {
	weight, height float64
}

func area(r Rectangle) float64  {

	return r.height*r.weight
}

func (h *Human) sayHi()  {
	fmt.Println("这个使用面向对象的思想来面对他们的数据分析",h.name)

}

func (s *Student) sayHi()  {

	fmt.Println("这个是来自学生的问候", s.name)
}



func main() {
	mark := Student{Human{"mark", 25, 200}, "computer Science"}
	fmt.Println("his name is ", mark.name)
	fmt.Println("his age is ", mark.age)
	fmt.Println("his srtt is ", mark.speciality)
	mark.age = 223
	fmt.Println(" mark changed his age and his age is ", mark.age)
	liliangbin, age := fallback(mark,mark)

	mark.sayHi()
	alpha := Human{"liliangib",22,333}
	alpha.sayHi()
	/*
	在go语言中面向对象的操作方法是改用另外的一个种继承的方式*

	*/
	fmt.Println("his name is ",liliangbin ,"his name fdhh ",age)
}
