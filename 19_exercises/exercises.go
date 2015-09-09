package main 
import "fmt"

func main () {
	var mySlice []interface {}
	mySlice = append(mySlice, 4)
	mySlice = append(mySlice, "string")
	mySlice = append(mySlice, 4.04)
	fmt.Println(mySlice)
	
	
	print ("this is a string")
	print (666)
	print (5.7894)
}

func print (x interface {}) {
	fmt.Println(x)
}