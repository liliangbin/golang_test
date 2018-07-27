package main

import (
	"html/template"
	"fmt"
	"net/http"
	"strings"
	"log"
	"golang_test/web/controller"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                    //这个地方一定要有这个函数，不然的话，他是不会把数据给弄出来的
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
	}
}
func hei(w http.ResponseWriter, r *http.Request) {

	controller.Hei(w, r)
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/hei", hei)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
