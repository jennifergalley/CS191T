package main
import "fmt"

func main () {
	myMap := map[string]string {
		"Jennifer":"Garner",
		"Alex":"Galley",
		"Alana":"Reed",
	}

	myMap["Navmit"] = "Danjeer"

	myMap["Navmit"] = "Something Else"

	delete(myMap, "Navmit")

	for key, val := range myMap {
		fmt.Println (key, " - ", val)
	}

	fmt.Println(len(myMap))

	if val, ok := myMap["Alex"]; ok {
		fmt.Println(val)
	}
}
