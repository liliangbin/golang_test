package main
import (
	"fmt"
)
type Book struct{
	book_id int
	book_nme string

}
//在于结构体这一块。我们需要重新定义一波 type
//slice 切片，动态数组，使用能够使用方法
func niubi  (num1 int,test string ) (int ){
	fmt.Println("nihao ")
	var book1  Book
	book1.book_id = 2
	book1.book_nme = "fdd"
	fmt.Println(book1.book_id)
	//切片  为空的时候 nil 
	var number = make([]int, 3,5)
	fmt.Printf("fff  %d \n",number[3])

	return 1
}

