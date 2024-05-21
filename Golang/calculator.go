package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Calculator function to perform operations based on the operand
func Calculator(value1, value2 float64, operand string) (float64, error) {
	// TODO: Implement Calculator function
	var result float64
	var err error
	switch operand {
	case "+":
		result = value1 + value2
	case "-":
		result = value1 - value2
	case "/":
		if value2 == 0 {
			err = fmt.Errorf("invalid value")
		} else {
			result = value1 / value2
		}
	case "*":
		result = value1 * value2
	default:
		err = fmt.Errorf("invalid case")
	}
	return result, err
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the expression (e.g., 10 + 20): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error reading input: %v\n", err)
		return
	}

	input = strings.TrimSpace(input)
	parts := strings.Split(input, " ")

	// TODO: Check if exact 3 parts are given. If not, throw an error
	if len(parts) != 3 {
		fmt.Println("enter a valid expression")
		return
	}
	// TODO: Take all 3 values and pass to calculator function
	value1, err := strconv.ParseFloat(parts[0], 64)
	value2, err := strconv.ParseFloat(parts[2], 64)

	// TODO: Print results
	result, err := Calculator(value1, value2, parts[1])
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Result: %v\n", result)
}
