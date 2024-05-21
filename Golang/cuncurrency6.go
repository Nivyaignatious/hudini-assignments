// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

package main
 
import "fmt"
 
func main() {
    messages := make(chan string)

	go func() {
		messages <- "hello"
	}()
	go func () {
		messages <- "world"
	}()
	for i:=0; i<2;i++{
	select{
    case message1 := <-messages:
	 fmt.Println(message1)
    case message2 := <-messages:
		fmt.Println(message2)
	}
  }
}

// Write a program that launches two goroutines. Each goroutine sends a series of messages to a channel.
// The main function should use a select statement to receive messages from both channels and print them.
// Hints:

// Use two separate channels.
// Use the select statement inside a loop to receive from the channels.

// -------------------------------------------------------------------------------------