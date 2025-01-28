package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 4))
}

func search(nums []int, target int) int {
	startIndex := 0
	endIndex := len(nums) - 1

	for startIndex <= endIndex {
		middleIndex := startIndex + (endIndex-startIndex)/2
		if nums[middleIndex] == target {
			return target
		} else if nums[middleIndex] < target {
			startIndex = middleIndex + 1
		} else if nums[middleIndex] > target {
			endIndex = middleIndex - 1
		}
	}
	return -1

}
