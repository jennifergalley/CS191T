package main
import "fmt"

func multThreads () {
	go func () {
		for i := 0; i < 10; i += 2 {
			fmt.Print(i)
		}
	}()
	
	go func () {
		for i := 0; i < 10; i += 3 {
			fmt.Print(i)
		}
	}()

	for i := 0; i < 10; i += 4 {
		fmt.Print(i)
	}
}

func channel () {
	ch := make (chan int, 5)
	
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println (<-ch)
	}
	
	close (ch)
}

func closedChannel () {
	ch := make (chan int, 5)
	
	for i := 0; i < 5; i++ {
		ch <- i
	}
	
	close (ch)
	
	for i := 0; i < 5; i++ {
		fmt.Println (<-ch)
	}
}

func deadlock () {
	ch := make (chan int, 5)
	fmt.Println (<-ch)
}


func main () {
	multThreads()
	
	channel ()
	
	closedChannel()
	
	deadlock()
}