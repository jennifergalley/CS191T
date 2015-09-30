package main 

import (
	"html/template"
	"log"
	"os"
)

func main () {
	name := "Jenny"
	
	//parse template
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	
	//execute template
	err = tpl.Execute(os.Stdout, name)
	if err != nil {
		log.Fatalln(err)
	}
}