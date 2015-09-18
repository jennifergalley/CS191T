package main

import (
	"fmt"
	"github.com/jennifergarner/CS191T/hello"
	"github.com/jennifergarner/CS191T/converters"
)

func main () {
	hello.Hello() 
	ft := converters.InchesToFeet (12.5)
	fmt.Println("12.5 inches is", ft, "feet")
}