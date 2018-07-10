package main

import (
	"fmt"
)

//这波主要是进行for循环的测试
func  forinit(){
	for a:=0;a<100;a++{
		if a%3==0 {
			fmt.Println("能够被3整除======>",a)
		} else if a%5==0 {
			fmt.Println("能够被5整除=====>",a)
		} else {
			fmt.Println("a 的值是多少 =====>",a)
		}
	}
}


func forarry(){
	var arr [10]int//初值为都为0
	for a:=0;a<10;a++{
		fmt.Println(arr[a])
	}
}
func maps()  {
	
	monthdays := map[string] int {
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, 
		}


		defer fmt.Println("hello world") //延迟输出，可以理解成一个栈结构，当数据最后搞定后再输出来


		for k,v:= range monthdays{ //对应的是一个j

			fmt.Println(k)
			fmt.Println(v)
			//函数作为值进行传递。我们使用的是她的额方法
			
		}

		a :=func ()  {
			
			fmt.Println("hhhh")
		}

		a()
}


func interesting(a int,b int )(x int ,y int)  {
	

	if a>b{
		return b,a 
	}else {
		return a ,b
	}
}
func main()  {
	//forinit();//输出对象
	//forarry();
/*
*arry 是一个实体的数据类型，但是slice却是一个引用类型。
* */

maps()

a,b :=interesting(2,5)
fmt.Println(a)
fmt.Println(b)

}