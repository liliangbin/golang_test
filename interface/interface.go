package main
import (
	"fmt"
)

type Human struct {
	name string
	age int 
	phone string 
}

type Student struct{
	Human
	school string 
	loan float32

}
type Employee struct{

	Human
	company string 
	money float32
}

func main(){
fmt.Println("hello world ")


}