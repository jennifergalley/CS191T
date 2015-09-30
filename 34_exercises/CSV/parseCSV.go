package main 

import (
	"html/template"
	"log"
	"net/http"
	"github.com/jennifergarner/CS191T/file"
)

func main () {
	data := file.ReadCSV("table.csv")
	
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		err = tpl.Execute(res, data)
		if err != nil {
			log.Fatalln(err)
		}
	})
	
	http.ListenAndServe(":9000", nil)
}