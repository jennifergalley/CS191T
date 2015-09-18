package main

import (
	"fmt"
	"encoding/json"
	"github.com/jennifergarner/CS191T/file" //to simplify file I/O
)

/*Create a struct, encode it as JSON, and write it out to stdout.*/
type Pasta struct {
	Noodle string
	Sauce string
	Cheese string
	Meat string
}

func encodeStruct () {
	myFav := Pasta{"Penne", "Red", "Parmesan", "None"}
	myJsonBytes, _ := json.Marshal(myFav)
	fmt.Println(string(myJsonBytes))
}

/*
Create a JSON file
Create a program that reads that JSON file into a struct, then writes it to stdout.
*/
func decodeJSON (filename string) {
	var obj Pasta
	contents := file.ReadAll (filename)
	json.Unmarshal([]byte(contents), &obj)
	fmt.Println(obj)
}

/*
Create a program which can read a csv file and write it as a JSON file.
*/
func csvToJSON (csvFile, jsonFile string) {
	//read CSV file
	data := file.ReadCSV(csvFile)
	
	//do stuff with data to get into struct
	var slice []Pasta
	for _, each := range data {
		var obj Pasta
		obj.Noodle = each[0]
		obj.Sauce = each[1]
		obj.Cheese = each[2]
		obj.Meat = each[3]
		slice = append (slice, obj)
	}
		
	//write to JSON file
	myJsonBytes, _ := json.Marshal(slice)
	file.WriteNew(jsonFile, string(myJsonBytes))
}

func main () {
	encodeStruct()
	decodeJSON("myFav.json")
	csvToJSON("myFav.csv", "myFav2.json")
}