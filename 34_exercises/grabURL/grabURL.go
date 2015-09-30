package main 

import (
	"html/template"
	"log"
	"net/http"
)

func main () {
	var url string
	
	//parse template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		url = req.RequestURI
		//execute template
		err = tpl.Execute(res, url)
		if err != nil {
			log.Fatalln(err)
		}
	})
	
	http.ListenAndServe(":9000", nil)
}