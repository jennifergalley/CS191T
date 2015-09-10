package main
import (
	"fmt"
	"strings"
	"math"
	"io/ioutil"
	"bufio"
	"os"
)

/*
create a program that
returns a map of the counts of each “word” in a string. (strings.Fields) 
WordCount(“test test”) → map[string]int{ “test”: 2 }
*/

func WordCount (s string) map[string]int {
	words := strings.Fields (s)
	result := make(map[string]int)
	
	for _, w := range words {
		result[w]++ //if key is not present, it returns the value's zero value, which is 0
	}
	
	return result
}


/* 
create a program that 
computes the average of a list of numbers, but removes the largest and smallest values. 
centeredAverage([]float64{1, 2, 3, 4, 100}) → 3
*/

func CenteredAverage (numbers []float64) float64 {
	if len(numbers) <= 2 {
		return 0.0
	}
	min := numbers[0]
	max := numbers[1]
	
	for _, val := range numbers {
		min = math.Min(min, val)
		max = math.Max(max, val)
	}
	
	sum := 0.0
	for _, val := range numbers {
		if val != max && val != min {
			sum += val
		}
	}
	
	return sum / float64(len(numbers) - 2)
}

/*
Write a program that can swap two integers 

x := 1 
y := 2
swap(&x, &y) 


should give you x=2 and y=1
*/

func swap (x, y *int) {
	temp := *x
	*x = *y
	*y = temp
	return
}

/*
Say that a "clump" in a list is a series of 2 or more adjacent elements of the same value. 
Return the number of clumps in the given list. 
countClumps([]int{1, 2, 2, 3, 4, 4}) → 2 
countClumps([]int{1, 1, 2, 1, 1}) → 2 
countClumps([]int{1, 1, 1, 1, 1}) → 1
*/

func countClumps (numbers []int) (count int) {
	clump := false
	for i, val := range numbers {
		if i > 0 {
			prev := numbers[i-1]
			if prev == val {
				clump = true
			} else if clump {
				clump = false
				count++
			}
		}
	}
	if clump {
		count++
	}
	return 
}

/*
Create your own version of cat which reads a file and dumps it to stdout.
*/
func cat (filename string) {
	contents, err := ioutil.ReadFile (filename)
	if err != nil {
		fmt.Println("error reading file")
	} else {
		fmt.Println(string(contents))	
	}
}

/*
Create a program which opens a file, reads a file, then writes the contents to a new file.
*/
func copy (oldFile string) {
	contents, err := ioutil.ReadFile (oldFile)
	if err != nil {
		fmt.Println("error reading file")
	} else {
		dot := strings.Index(oldFile, ".")
		name := oldFile[:dot]
		ext := oldFile[dot:]
		newFile := name + "-copy" + ext
		ioutil.WriteFile(newFile, contents, 0777)	
	}
}

/*
Create your own version of cp which reads a file and writes it to another file.
*/
func cp (oldFile, newFile string) {
	contents, err := ioutil.ReadFile (oldFile)
	if err != nil {
		fmt.Println("error reading file")
	} else {
		ioutil.WriteFile(newFile, contents, 0777)	
	}
}

/*
Q: Why would you use an embedded anonymous unnamed field instead of a normal named field?
A: So you could access the anonymous field's properties as though they belonged to the outer object.
*/

/*
Create a program which converts the first character of each line in a file to uppercase and writes it to stdout.
*/
func capitalizeLine (filename string) {
	file, _ := os.Open(filename)
    fscanner := bufio.NewScanner(file)
    for fscanner.Scan() {
		line := fscanner.Text()
		first := line[0]
		newFirst := strings.ToUpper(string(first))
		newline := newFirst + line[1:]
        fmt.Println(newline)
    }
}

/*
Create a program which capitalizes the first letter of every word from a text file and writes it to stdout.
*/
func capitalizeWords (filename string) {
	file, _ := os.Open(filename)
    fscanner := bufio.NewScanner(file)
    for fscanner.Scan() {
		line := fscanner.Text()
		words := strings.Fields(line)
		for _, w := range words {
			newFirst := strings.ToUpper(string(w[0]))
			newWord := newFirst + w[1:]
			fmt.Print(newWord, " ")
		}
		fmt.Println()
    }
}

/*
Create a program which capitalizes every other word (capitalizes the entire word) from a text file and writes it to stdout.
*/
func capitalizeOddWords (filename string) {
	file, _ := os.Open(filename)
    fscanner := bufio.NewScanner(file)
	odd := true
    for fscanner.Scan() {
		line := fscanner.Text()
		words := strings.Fields(line)
		for _, w := range words {
			if odd {
				fmt.Print(strings.ToUpper(w), " ")
				odd = false
			} else {
				fmt.Print(w, " ")
				odd = true
			}
		}
		fmt.Println()
    }
}



func main () {
	fmt.Println("Word Count")
	fmt.Println(WordCount("test test"))
	fmt.Println()
	
	fmt.Println("Centered Average")
	fmt.Println(CenteredAverage([]float64{1, 2, 3, 4, 100}))
	fmt.Println()
	
	fmt.Println("Swap")
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Println(x, y)
	fmt.Println()

	fmt.Println("Count Clumps")
	fmt.Println(countClumps([]int{1, 2, 2, 3, 4, 4}))
	fmt.Println(countClumps([]int{1, 1, 2, 1, 1}))
	fmt.Println(countClumps([]int{1, 1, 1, 1, 1}))
	fmt.Println()

	fmt.Println("Cat")
	cat("test.txt")
	fmt.Println()

	fmt.Println("Copy")
	copy("test.txt")
	cat("test-copy.txt")
	fmt.Println()	
	
	fmt.Println("cp")
	cp("test.txt", "newFile.txt")
	cat("newFile.txt")
	fmt.Println()
	
	fmt.Println("capitalizeLine")
	capitalizeLine("test.txt")
	fmt.Println()
	
	fmt.Println("capitalizeWords")
	capitalizeWords("test.txt")
	fmt.Println()
	
	fmt.Println("capitalizeOddWords")
	capitalizeOddWords("test.txt")
	fmt.Println()
}