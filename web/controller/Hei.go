package controller

import (
	"net/http"
	"fmt"
	"html/template"
	"golang_test/web/model"
	"golang_test/web/service"
)

func Hei(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" {

		t, _ := template.ParseFiles("index.html")
		fmt.Println("index")
		t.Execute(w, nil)

	} else {
		var user model.User
		user.Name = r.Form["name"][0]
		user.Info = r.Form["info"][0]

		hh := service.Add(&user)
		fmt.Println(r.Form["name"], "indexindex")
		fmt.Fprint(w, hh)

	}
}
