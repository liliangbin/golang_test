
package main

import (
	"fmt"
)

func main()  {
	
	var countryCapital map[string]string
	countryCapital = make(map[string]string)
	countryCapital["france"] = "parise"
	countryCapital["america"] = "newyour"//这个地方是把数据添加进图里面去

	for country :=range countryCapital{
		fmt.Println(country,"首都是",countryCapital[country])	}

	capitlal,ok :=countryCapital["meiguo"]
	if ok {
		fmt.Println("美国的首都是",capitlal)
	}else{
		fmt.Println("没有查找到他的首度")
	}

}