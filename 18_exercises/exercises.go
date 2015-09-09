package main 
import "fmt"

type Person struct {
	name string
	age int
}

type Programmer struct {
	Person
	language string
	os string
	company string
}

func (p Programmer) skills () {
	fmt.Println (p.name + " is good at " + p.language + " on " + p.os + " and works at " + p.company)
}

func main () {
	jennifer := Programmer {
		Person{
			name: "Jennifer",
			age: 21,
		},
		"C++",
		"Windows",
		"Microsoft",
	}
	
	jennifer.skills()
}