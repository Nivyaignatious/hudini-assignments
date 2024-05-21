// Objective:
// Use sync.WaitGroup to wait for multiple goroutines to complete.

// Task:

package main
import "fmt"
import "sync"


func differentMsg(msg string,wg *sync.WaitGroup){
	fmt.Println(msg)
	wg.Done()
}

func main( ){
	// create a waitgroup
	var wg sync.WaitGroup
	wg.Add(3)
	// call three go routines
	go differentMsg("hell",&wg)
	go differentMsg("go",&wg)
	go differentMsg("away",&wg)
	
	// wait for three go routines
	wg.Wait()
}
 

// Write a program that launches three goroutines, each printing a different message.
// Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.

// -------------------------------------------------------------------------------------