package main

import (
	"html/template"
	"log"
	"net/http"
)

type Person struct {
	FirstName string
	LastName string
}

func main () {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fname := req.FormValue("first")
		lname := req.FormValue("last")
		
		err := tpl.Execute(res, Person{fname, lname})
		if err != nil {
			http.Error(res, err.Error(), 500)
			log.Fatalln(err)
		}
	})
	
	http.ListenAndServe(":9000", nil)
	
}