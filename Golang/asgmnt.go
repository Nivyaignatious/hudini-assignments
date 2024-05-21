package main

import "fmt"

// question 1
func avg(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	return sum / len(numbers)
}

// question 2
func checkage(age int) string {
	if age > 150 {
		return "exceeding age limit"
	} else if age < 0 {
		return "negative age limt"
	} else if age < 18 {
		return "minor"
	} else if age > 18 && age < 35 {
		return "youth"
	} else {
		return "adult"
	}
}

// question 3
func reverse(str string) string {
	var strg string = ""
	for i := len(str) - 1; i >= 0; i-- {
		strg += string(str[i])
	}
	return strg
}

// question 4
func largest(array []int) int {
	max := 0
	for i := 0; i < len(array); i++ {
		if array[i] > max {
			max += array[i]
		}
	}
	return max
}

// question 5
type Counter struct {
	count int
}

func (counter Counter) add() Counter {
	counter.count++
	return counter
}
func (counter Counter) sub() Counter {
	counter.count--
	return counter
}
func (counter Counter) display() Counter {
	fmt.Println(counter)
	return counter
}
func (counter Counter) reset() Counter {
	counter
}

func main() {
	//1
	s := []int{10, 23, 45, 6, 9}
	fmt.Println(avg(s))
	//2
	ages := 25
	fmt.Println(checkage(ages))
	//3
	fmt.Println(reverse("hello"))
	//4
	fmt.Println(largest(s))
	//5
	counter := Counter{count: 0}
	counter = counter.add()
	counter = counter.add()
	counter.display()
	counter = counter.sub()
	counter.display()

}
