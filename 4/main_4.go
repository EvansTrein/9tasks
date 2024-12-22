package main

import (
	"fmt"
)

// сложность O(n)
func SetDifference(slice1, slice2 []string) (diffSlice []string) {
	counterSlice2 := make(map[string]int)

	for _, el := range slice2 {
		counterSlice2[el] += 1
	}

	for _, el := range slice1 {
		if count := counterSlice2[el]; count == 0 {
			diffSlice = append(diffSlice, el)
		}
	}

	return diffSlice
}

func main() {
	fmt.Println(SetDifference([]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}, []string{"banana", "date", "fig"}))
	fmt.Println(SetDifference([]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1", "one", "two"}, []string{"banana", "date", "fig", "two"}))
}
