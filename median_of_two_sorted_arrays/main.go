package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello world")
	nums1 := make([]int, 0)
	nums2 := make([]int, 0)

	findMedianSortedArrays(nums1, nums2)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    // TODO try with two pointers
	arrLengths := len(nums1) + len(nums2)

	// what happens if it not odd
	isOdd := arrLengths%2 == 1

	newArr := append(nums1, nums2...)

	sortedArr := sort.IntSlice(newArr)

	sortedArr.Sort()
	medianIndex := arrLengths / 2

	if !isOdd {
		return float64(sortedArr[medianIndex])
	}

	res := float64(sortedArr[medianIndex] + sortedArr[medianIndex+1]) / 2.0

	return res
}
