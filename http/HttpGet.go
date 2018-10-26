package main


import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main() {

	resp, err := http.Get("http://jwc.upc.edu.cn/")

	if err != nil {

		panic(err.Error())
	}
	defer resp.Body.Close()


	fmt.Println(resp.Body)
	body ,err :=ioutil.ReadAll(resp.Body)

	if err!=nil {

		panic(err.Error())
	}
	htmlBody :=string(body)
	fmt.Println(htmlBody)


	fmt.Println(string(body))

	err =ioutil.WriteFile("inde.html",[]byte(htmlBody),0633)

	if err != nil {
		panic(err.Error())
	}
}
