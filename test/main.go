package main

import (
	"fmt"
)

func isAnagram(s string, t string) bool {
	return false
}

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string)
	for _, v := range strs {
		is := false
		for key, _ := range resultMap {
			if isAnagram(v, key) {
			append(resultMap[key], v)
				is = true
				break
			}
		}
		if !is {
			resultMap[v] = make([]string, 20)
			append(resultMap[v], v)
		}
	}
	return make([][]string, 20)
}

func main() {

	fmt.Printf("zort")
}
