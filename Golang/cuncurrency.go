// Objective:
// Learn how to send and receive values using channels.

// Task:

// Write a program that creates a goroutine to send a message "Hello, World!" to a channel.
// The main function should receive the message from the channel and print it.
// Hints:
package main

import "fmt"

func printInChannel(messages chan string) {
	messages <- "hello world"
}

func main() {
	ch := make(chan string)
	go printInChannel(ch)
	msg := <-ch
	fmt.Println(msg)
}

// Use an unbuffered channel for simplicity.

// -------------------------------------------------------------------------------------



