// question 1
//In this example you have to validate if a user input string is alphanumeric. The given string is not nil/null/NULL/None, so you don't have to check that.

//The string has the following conditions to be alphanumeric:

//At least one character ("" is not valid)
//Allowed characters are uppercase / lowercase latin letters and digits from 0 to 9
//No whitespaces / underscore
package kata

func alphanumeric(str string) bool {
  if len(str) == 0 {
        return false
    }
    for _, i := range str {
        if !((i >= 'a' && i <= 'z') || (i >= 'A' && i <= 'Z') || (i >= '0' && i <= '9')) {
            return false
        }
    }
    return true
}

//question 2
// Write an algorithm that takes an array and moves all of the zeros to the end, preserving the order of the other elements.

//MoveZeros([]int{1, 2, 0, 1, 0, 1, 0, 3, 0, 1}) // returns []int{ 1, 2, 1, 1, 3, 1, 0, 0, 0, 0 }
package kata

func MoveZeros(arr []int) []int {
  count := 0 
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[count] = arr[i]
			count++
		}
	}
	for count < len(arr) {
		arr[count] = 0
		count++
	}
  return arr
}

//question 3
//Your task is to write a function which returns the sum of a sequence of integers.

//The sequence is defined by 3 non-negative values: begin, end, step.

//If begin value is greater than the end, your function should return 0. If end is not the result of an integer number of steps, then don't add it to the sum. See the 4th example below.

//Examples

//2,2,2 --> 2
2,6,2 --> 12 (2 + 4 + 6)
1,5,1 --> 15 (1 + 2 + 3 + 4 + 5)
1,5,3  --> 5 (1 + 4)

package kata


func SequenceSum(start, end, step int) int {
  sum:=0
  if start>end {
    return 0
  }
  if start== end{
    return end
  }
  if start<end{
    for i:=start; i<=end; i+=step{
    sum +=i
    }
    }
  return sum
}

// question4
//In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

//Examples
//HighAndLow("1 2 3 4 5")  // return "5 1"

package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	highest:= -1000
  lowest:= 1000
	for _, numStr := range strings.Fields(in) {
		num, _ := strconv.Atoi(numStr)
		if num > highest {
			highest = num
		}
		if num < lowest {
			lowest = num
		}
	}
	return fmt.Sprintf("%d %d", highest, lowest)
}

//question5
//You are going to be given an array of integers. Your job is to take that array and find an index N where the sum of the integers to the left of N is equal to the sum of the integers to the right of N.

//If there is no index that would make this happen, return -1.
package kata

func FindEvenIndex(arr []int) int {

	for i := 0; i < len(arr); i++ {
		leftSum := 0
		for j := 0; j <= i; j++ {
			leftSum += arr[j]
		}

	rightSum := 0
		for j := i; j < len(arr); j++ {
			rightSum += arr[j]
		}   
		if leftSum == rightSum {
			return i
		}
  }
	return -1
}