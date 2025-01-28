package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Hello, World!")
}

func merge(nums1 []int, m int, nums2 []int, n int) {
    for n > 0 {
        nums1[m] = nums2[n-1]
        m++
        n--
    }

    sort.Ints(nums1)
}
