package main

import "fmt"

func main() {
	a := "ab"
	b := "eidbaooo"
	res := checkInclusion(a, b)
	fmt.Println(res)
}

func areMapsEqual(map1, map2 map[rune]int) bool {

	for ch, val := range map1 {
		if val != map2[ch] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {

	if len(s1) > len(s2) {
		return false
	}

	windowLength := len(s1)
	s1Map := make(map[rune]int)
	s2Map := make(map[rune]int)
	for _, ch := range s1 {
		s1Map[ch]++
	}

	for i := 0; i < windowLength; i++ {
		s2Map[rune(s2[i])]++
	}
	if areMapsEqual(s1Map, s2Map) {
		return true
	}

	for i := 0; i < len(s2)-windowLength; i++ {
		s2Map[rune(s2[i])]--
		if s2Map[rune(s2[i])] == 0 {
			delete(s2Map, rune(s2[i]))
		}
		s2Map[rune(s2[i+windowLength])]++

		if areMapsEqual(s1Map, s2Map) {
			return true
		}

	}

	return false
}
