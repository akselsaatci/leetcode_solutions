package main

import (
	"math"
	"sort"
)

func main() {

}

func isValidSudoku(board [][]byte) bool {

	return true
}

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	for _, val := range strs {
		s := []rune(val)
		sort.Slice(s, func(i int, j int) bool { return s[i] < s[j] })
		anagramMap[string(s)] = append(anagramMap[string(s)], val)
	}
	result := [][]string{}

	for _, val := range anagramMap {
		result = append(result, val)

	}

	return result
}

func containsNearbyDuplicate(nums []int, k int) bool {
	indexMap := make(map[int]int)

	for i, num := range nums {
		if lastIndex, exists := indexMap[num]; exists {
			if i-lastIndex <= k {
				return true
			}
		}
		indexMap[num] = i
	}

	return false
}

func intersect2(nums1 []int, nums2 []int) []int {
	table := make(map[int]struct{})

	for _, num := range nums1 {
		table[num] = struct{}{}
	}

	var result []int
	for _, num := range nums2 {
		if _, found := table[num]; found {
			result = append(result, num)

		}
	}
	resultMap := make(map[int]*struct{})

	for _, val := range result {
		resultMap[val] = new(struct{})
	}
	result = []int{}
	for key := range resultMap {
		result = append(result, key)
	}

	return result
}

func isIsomorphic(str1 string, str2 string) bool {

	isoMap1 := make(map[rune]rune)
	isoMap2 := make(map[rune]rune)
	if len(str1) != len(str2) {
		return false
	}

	for index, value := range str1 {
		mapValue, exists := isoMap1[value]
		if !exists {
			isoMap1[value] = rune(str2[index])
			continue
		}
		if mapValue != rune(str2[index]) {
			return false
		}
	}

	for index, value := range str2 {
		mapValue, exists := isoMap2[value]
		if !exists {
			isoMap2[value] = rune(str1[index])
			continue
		}
		if mapValue != rune(str1[index]) {
			return false
		}
	}

	return true
}

func intersect(nums1 []int, nums2 []int) []int {

	return []int{}
}

func firstUniqChar(s string) int {
	charMap := make(map[rune][]int)
	var order []rune

	for index, value := range s {
		if _, exists := charMap[value]; !exists {
			order = append(order, value)
		}
		charMap[value] = append(charMap[value], index)
	}

	for _, key := range order {
		if len(charMap[key]) == 1 {
			return charMap[key][0]
		}
	}

	return -1
}
func findRestaurant(list1 []string, list2 []string) []string {

	zortMap := make(map[string]int)

	for index, value := range list1 {
		zortMap[value] = index

	}

	result := make([]string, 0)
	minSum := math.MaxInt32
	for index, value := range list2 {
		if mapValue, ok := zortMap[value]; ok {
			if mapValue+index < minSum {
				result = []string{value}
				minSum = mapValue + index
			} else if mapValue+index == minSum {
				result = append(result, value)
			}
		}
	}
	return result

}

func twoSum(nums []int, target int) []int {

	table := make(map[int]int)

	for index, value := range nums {
		temp := target - value
		mapVal, ok := table[temp]

		if ok {
			return []int{mapVal, index}
		}
		table[value] = index
	}
	return make([]int, 0)

}
func nextNumberCycle(n int) int {

	newNumber := 0

	for n != 0 {
		temp := n % 10
		println("mod ", temp)
		newNumber = temp*temp + newNumber
		n = n / 10
	}
	println(newNumber)

	return newNumber
}

func isHappyCycle(n int) bool {
	slow := n
	fast := nextNumberCycle(n)

	// First loop to find the cycle
	for fast != 1 && slow != fast {
		slow = nextNumberCycle(slow)
		fast = nextNumberCycle(nextNumberCycle(fast))
	}

	// If fast reaches 1, it is a happy number
	return fast == 1
}
func nextNumber(n int) int {

	newNumber := 0

	for n != 0 {
		temp := n % 10
		println("mod ", temp)
		newNumber = temp*temp + newNumber
		n = n / 10
	}
	println(newNumber)

	return newNumber
}

func isHappy(n int) bool {

	seenMap := make(map[int]*struct{})
	for seenMap[n] == nil {
		if n == 1 {
			return true
		}
		seenMap[n] = new(struct{})
		n = nextNumber(n)

	}

	return false
}

func singleNumber(nums []int) int {
	table := make(map[int]*struct{})

	for _, value := range nums {
		if table[value] != nil {
			table[value] = nil
			println(value)
			println(table[value])
			continue
		}
		table[value] = new(struct{})

	}
	for key, value := range table {
		if value != nil {
			return key
		}
	}
	return 0
}
