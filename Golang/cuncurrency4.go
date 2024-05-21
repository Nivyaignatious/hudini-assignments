// Objective:
// Understand how to use channels for communication between goroutines.

// Task:

package main

import "fmt"

func numbers(messages chan int) {
	for i := 1; i <= 10; i++ {
		messages <- i
	}
	close(messages)
}
func main() {
	messages := make(chan int)
	go numbers(messages)
	for i := range messages {
		fmt.Println(i)
	}
}

// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:

// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.

// -------------------------------------------------------------------------------------
