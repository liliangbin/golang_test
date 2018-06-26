package main

import "fmt"
var a = "cai niao  jiaocheng  "
var b string = "runbooob.com"
var c bool
var (

	alpha int 
	beta string 
)

func main(){
	lilangbin := "intet"
	const LENGTH int = 10;
	
	//var ttt = 12 
	fmt.Println("hello world \n ")
	fmt.Println("在go 语言中，：= 这种形式的赋值语句只能是用于函数体里面    " , lilangbin, len(lilangbin)   )
	fmt.Println(len(lilangbin))
	// 这个时候还是可以像字符穿的添加一样使用
	//在变量的声明中全局变量是可以声明但是不使用的，但是局部变量是不允许声明不适用的
	//条件判断语句
	if len(lilangbin) < 20{
		fmt.Println("数据的迁移问题")
	} else {
		fmt.Println(lilangbin)
	}
	//for循环开始简化了
	for index := 0; index < 15; index++ {
		
		fmt.Println(lilangbin)
	}
	var ret, test = max(2,3)
	fmt.Println(ret,test)
}

func max(num1, num2 int) (int,string ) {
	/* 定义局部变量 */
	var result int
 
	if (num1 > num2) {
	   result = num1
	} else {
	   result = num2
	}
	return result ,"liliangbin"
 }



