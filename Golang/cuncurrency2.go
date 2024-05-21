// Objective:
// Learn how to create and use goroutines.
 
// Task:
package main
import "fmt"
import "sync"
import "time"


func display(wg *sync.WaitGroup){
	for i:=1; i<=5; i++{
		fmt.Println(i)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func main( ){
	// create a waitgroup
	var wg sync.WaitGroup
	// call three go routines
	for i:=0; i<=3; i++{
		wg.Add(1)
       go display(&wg)
	}
	// wait for three go routines
	wg.Wait()
}
 
// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Hints:
 
// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.
 // -------------------------------------------------------------------------------------
 