package main

import "fmt"

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(input))
}

func maxArea(height []int) int {
	maxArea := -1
	arrayLen := len(height)

	i := 0
	j := arrayLen - 1
	for {
		localMax := min(height[i], height[j]) * (j - i)
		fmt.Println(localMax)
		if localMax > maxArea {
			maxArea = localMax
		}
		if height[i] <= height[j] {
			i++
		} else {
			j--
		}
		if i == j {
			break
		}
	}

	return maxArea
}
func min(x, y int) int {
	if x <= y {
		return x
	} else {
		return y
	}

}
