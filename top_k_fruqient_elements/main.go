package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("ZortingenZortingen")
	var f_test, s_test []int
	f_test = []int{2, 2, 3, 1, 1, 1}
	s_test = []int{1}
	fmt.Println(topKFrequent(f_test, 2))
	fmt.Println(topKFrequent(s_test, 1))

}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)

	for _, val := range nums {
		countMap[val]++
	}

	var sorted []int
	for key := range countMap {
		sorted = append(sorted, key)
	}

	sort.Slice(sorted, func(i, j int) bool {
		return countMap[sorted[i]] > countMap[sorted[j]]
	})

	resNums := make([]int, k)
	copy(resNums, sorted[:k])

	return resNums
}
