package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(myAtoi("20000000000000000000"))
}

//Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer.
//
//The algorithm for myAtoi(string s) is as follows:
//
//Whitespace: Ignore any leading whitespace (" ").
//Signedness: Determine the sign by checking if the next character is '-' or '+', assuming positivity if neither present.
//Conversion: Read the integer by skipping leading zeros until a non-digit character is encountered or the end of the string is reached. If no digits were read, then the result is 0.
//Rounding: If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then round the integer to remain in the range. Specifically, integers less than -231 should be rounded to -231, and integers greater than 231 - 1 should be rounded to 231 - 1.
//Return the integer as the final result.

func isNumeric(c rune) bool {
	return c >= 48 && c <= 57
}

func myAtoi(s string) int {

	isItLeading := true
	isNegative := false
	var charTable []rune

	for _, v := range s {

		if v == ' ' && isItLeading {
			continue
		}

		if isItLeading && (v == '-') {
			isNegative = true
			isItLeading = false
			continue
		}
		if isItLeading && (v == '+') {
			isItLeading = false
			continue
		}

		if !isNumeric(v) {
			break
		}
		isItLeading = false
		charTable = append(charTable, v)

	}

	charTableLen := len(charTable)

	res := 0
	for i, v := range charTable {
		if res >= math.MaxInt32 {
			fmt.Println("MAXINTT")
			if isNegative {
				return -2147483647
			}
			return 2147483647
		}
		fmt.Println(i, v)
		numValue := int(v) - 48
		currentNum := float64(numValue) * math.Pow(10, float64(charTableLen-i-1))
		res += int(currentNum)

	}
	if isNegative {
		res = -res
	}

	return res
}
