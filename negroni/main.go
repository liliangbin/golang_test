package main

import (
	"net/http"
	"fmt"
	"github.com/codegangsta/negroni"
)

func main()  {

	mux :=http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer,"welcome to newyork ")
	})

	n:=negroni.Classic()
	n.Use(negroni.NewLogger())
	n.UseHandler(mux)
	n.Run(":3000")

}
