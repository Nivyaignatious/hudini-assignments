package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum - 5
	y = x * sum
	return
}

func main() {
	fmt.Println(split(10))
}
