
package main

import "fmt"

var liliangbin int = 64 
type In interface{

Get() int 
Set(int) 
}

type Student struct{
	Age  int
}

func (s Student) Get() int{
	return s.Age
}

//我们使用的方法是不同他的
func (s *Student) Set(age int) {
	s.Age = age
}

//指针类型。
//函数的调用，以及方法的实现对于地址的调用是不同的。
func f(in In){
	in.Set(10)
	fmt.Println(in.Get())
}
//

func printAll(vals []interface{}){
	for _,val :=range vals{
		fmt.Println(val)
	}
}
func main() {

	fmt.Println("hello world ")
	s :=Student{}
	f(&s)
	fmt.Println(&s.Age)

//	name := []string{"strnag","dvid","csar"}

}