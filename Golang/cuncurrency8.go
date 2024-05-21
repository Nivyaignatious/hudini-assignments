// Objective:
// Learn how to use the select statement to handle multiple channels.

// Task:

package main

import (
    "fmt"
)

func main() {

    c1 := make(chan string)
    c2 := make(chan string)

    go func() {
        c1 <- "one"
    }()
    go func() {
        c2 <- "two"
    }()

            fmt.Println("received", msg1)

            fmt.Println("received", msg2)
        }
    

// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.