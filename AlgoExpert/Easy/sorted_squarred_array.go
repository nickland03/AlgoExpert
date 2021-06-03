package main

import (
	"sort"
)

func SortedSquaredArray(array []int) []int {
	var squaredArray []int

	for i := len(array)-1; i >= 0; i-- {
		squaredNum := array[i] * array[i]
		squaredArray = append(squaredArray, squaredNum)
	}

	sort.Ints(squaredArray)
	return squaredArray
}