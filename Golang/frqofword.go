package main

import (
	"bufio"
  "fmt"
  "strings"
  "os"
)


func freq(text string){
	input := strings.TrimSpace(text)
	
	str := strings.Split(input," ")
	
	frequency := make(map[string]int)
	for _ , words := range str{
		 frequency[words]++
	}
	 for words, count := range frequency {
		fmt.Println(words,count)
	  }
}
func main (){
	fmt.Println("Enter the block of code:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	freq(input)
	
}