package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	tial := s.Len() - 1
	for head := 0; head < s.Len()/2; head++ {
		if s.Less(head, tial) {
			return false
		}
	}
	return true
}

var testStrings1 = []string{"abc", "abcddcba", "abc"}
var testStrings2 = []string{"abc", "abcddcba", "123"}
var testStrings3 = []string{"1", "2", "3"}
var testStrings4 = []string{}
var testStrings5 = []string{"abc"}

type stringSlice []string

func (s stringSlice) Len() int {
	return len(s)
}

func (s stringSlice) Less(i, j int) bool {
	return s[i] < s[j] || s[i] > s[j]
}

func (s stringSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	fmt.Println(IsPalindrome(stringSlice(testStrings1)))
	fmt.Println(IsPalindrome(stringSlice(testStrings2)))
	fmt.Println(IsPalindrome(stringSlice(testStrings3)))
	fmt.Println(IsPalindrome(stringSlice(testStrings4)))
	fmt.Println(IsPalindrome(stringSlice(testStrings5)))

}
