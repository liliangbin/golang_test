package main

import (
	"fmt"
)

type Node struct{

	i int
	data [10]int
} 

func (s *Node) push(k int)  {
	if s.i +1 >9  {
		return
	}
	s.data[s.i] = k
	s.i++
}

func (s *Node) pop()(a int)  {
	s.i--
	if s.i>=0 {
		return s.data[s.i-1]
		
	}
	return 
}
func maopao(in  []int) {
	
	for i:=0 ;i<len(in);i++{
		for j:=i+1;j<len(in);j++{
			if in[j]<in[i] {
				in[i],in[j] = in[j],in[i]
			}
		}
	}
}

func reason(mapp map[string]string)  {
	

	mapp["sdfasdfas"]= "asdfa"
	mapp["fasdfsa"] = "asdfasf" 
	defer fmt.Println("构造函树")
}
func main()  {
	var node Node
	node.push(25)
	node.push(33) //面向对象的思想来处理数据。
	fmt.Printf("node  %v",node)

//假设我们使用面向过程的思想来做。不是很理解。

//冒泡排序
in := []int{2,4,5,3,1,8} //slice是一个应用类型，所以不用返回。
maopao(in)
fmt.Println(in)
mapdd := make(map[string]string)

mapdd["test"] = "dsadfasd"
reason(mapdd)
fmt.Println(mapdd)

}