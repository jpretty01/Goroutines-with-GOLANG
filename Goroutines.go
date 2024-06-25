package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for i := 'A'; i <= 'E'; i++ {
		fmt.Printf("%c\n", i)
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	go printNumbers() // Start a new goroutine
	go printLetters() // Start another goroutine

	// Sleep the main function to wait for the goroutines to complete
	// This is only for demonstration purposes; typically, you would use synchronization primitives like WaitGroups.
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}
