package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	substrs := strings.Fields(s)
	fmt.Println(substrs)

	// key, value => word, count
	// iterate over substrs, if key in map, value++, else create

	counts := make(map[string]int)

	for _, word := range substrs {
		_, ok := counts[word]

		if ok {
			counts[word] += 1
		} else {
			counts[word] += 1
		}
	}

	return counts
}

func main() {
	fmt.Println(WordCount("This exercise from the golang tour This tour is very very very ... "))
}
