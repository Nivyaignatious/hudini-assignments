// Objective:
// Learn how to use a buffered channel.
 
// Task:
package main

import "fmt"

func integers(msg chan int){
	for i:=1;i<=3; i++{
    msg <- i  
    }
	close(msg)
}
func main(){
	msg:= make(chan int,3)
	go integers(msg)
	for i:= range msg {
		fmt.Println(i)
	}

}
 // Write a program that creates a buffered channel with a capacity of 3.
// The main function should send three integers (1, 2, 3) to the buffered channel without blocking.
// Then, receive and print the integers from the channel.
 
// -------------------------------------------------------------------------------------