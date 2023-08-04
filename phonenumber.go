// https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
package main

import "fmt"

var chars = [][]string{
	{"a", "b", "c"},
	{"d", "e", "f"},
	{"g", "h", "i"},
	{"j", "k", "l"},
	{"m", "n", "o"},
	{"p", "q", "r", "s"},
	{"t", "u", "v"},
	{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	numCombinations := 1
	for _, d := range digits {
		if d != '9' && d != '7' {
			numCombinations *= 3
		} else {
			numCombinations *= 4
		}
	}
	fmt.Println(numCombinations)
	combinations := make([]string, numCombinations)
	currentNumCombinations := numCombinations
	for _, d := range digits {
		c := chars[d-'2']
		currentNumCombinations /= len(c)
		for i := 0; i < numCombinations; i++ {
			combinations[i] += c[(i/currentNumCombinations)%len(c)]
		}
	}

	return combinations
}
