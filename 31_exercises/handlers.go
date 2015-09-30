package main 

import (
	"io"
	"net/http"
)

type DogHandler int 
func (h DogHandler) ServeHTTP (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, "<img src='http://images4.fanpop.com/image/photos/15800000/Siberian-Husky-Puppy-puppies-15897208-1280-800.jpg' />")
}

type CatHandler int 
func (h CatHandler) ServeHTTP (res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(res, "<img src='http://www.brendasinclairauthor.com/wp-content/uploads/2012/02/kitten.png' />")
}

func main () {
	var dog DogHandler
	var cat CatHandler
	
	mux := http.NewServeMux()
	mux.Handle("/dog/", dog)
	mux.Handle("/cat/", cat)
	
	http.ListenAndServe (":9000", mux)
}