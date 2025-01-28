package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(input))
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i, v := range nums {
		p1 := i + 1
		p2 := len(nums) - 1
		for {
			if p1 >= p2 || p2 == i || p1 == i {
				break
			}
			if nums[p1]+nums[p2]+v == 0 {
				res = append(res, []int{v, nums[p1], nums[p2]})
				p1 += 1
			}

			if nums[p1]+nums[p2]+v < 0 {
				p1 += 1
				continue
			}

			if nums[p1]+nums[p2]+v > 0 {
				p2 -= 1
				continue
			}
		}

	}
	return res
}
