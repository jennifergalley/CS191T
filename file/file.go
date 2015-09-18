package file

import (
	"os"
	"encoding/csv"
	"io/ioutil"
)

/*
Read a file, returning all contents as a string.
*/
func ReadAll (filename string) string {
	contents, _ := ioutil.ReadFile (filename)
	return string(contents)
}

/*
Create a new file with name and extension filename, writing string contents to it.
*/
func WriteNew (filename, contents string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	if _, err = f.WriteString(contents); err != nil {
		panic(err)
	}
}

/*
Read a CSV file, returning a [][]string of the contents for each line, and each field
*/
func ReadCSV (filename string) [][]string {
	f, _ := os.Open(filename)
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	data, _ := reader.ReadAll()
	return data
}