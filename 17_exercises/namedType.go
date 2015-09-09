package main 
import "fmt"

type Programmer struct {
	name string
	language string
	os string
	company string
}

func (p Programmer) skills () {
	fmt.Println (p.name + " is good at " + p.language + " on " + p.os + " and works at " + p.company)
}

func main () {
	jennifer := Programmer {
		name: "Jennifer",
		language: "C++",
		os: "Windows",
		company: "Microsoft",
	}
	
	jennifer.skills()
}