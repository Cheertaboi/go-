package main

import (
	"fmt"
	"sync"
)

/*func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)

}*/

var msg string

func updateMessage(s string) {
	msg = s
}

func printMessage() {
	gg.Done()
	fmt.Println(msg)
}

var gg sync.WaitGroup

func main() {
	/*var wg sync.WaitGroup
	// if you run the program with this line uncommented, and the lines 20 commented,
	// everything works as expected
	//go printSomething("This is the first thing to be printed!")
	words := []string{"alpha", "gama", "delta", "version", "five", "six", "seven"}
	//integer := [...]int{1, 2, 3, 4, 5, 65, 67, 7, 8, 0}
	wg.Add(7)
	for i, x := range words {
		//	fmt.Printf("%d", i)
		go printSomething(fmt.Sprint(i, x), &wg)
	}

	wg.Wait()
	wg.Add(1)

	printSomething("This is the second thing to be printed!", &wg)
	*/

	msg = "Hello, world!"
	gg.Add(1)
	go updateMessage("Hello, universe!")
	printMessage()
	gg.Wait()
	gg.Add(1)
	go updateMessage("Hello, cosmos!")
	printMessage()
	gg.Wait()
	gg.Add(1)
	go updateMessage("Hello, world!")

	printMessage()
	gg.Wait()
}
