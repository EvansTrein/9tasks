package main

import "fmt"

// сложность O(n)
func SetIntersection(slice1, slice2 []int) (isInter bool, interSet []int) {
	counter := make(map[int]int, len(slice1))

	for _, el := range slice1 {
		counter[el] += 1
	}

	for _, el := range slice2 {
		if counter[el] > 0 {
			interSet = append(interSet, el)
			isInter = true
		}
	}

	return isInter, interSet
}

func main() {
	fmt.Println(SetIntersection([]int{65, 3, 58, 678, 64}, []int{64, 2, 3, 43}))
}