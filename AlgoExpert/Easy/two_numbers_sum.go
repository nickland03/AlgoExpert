package main

func TwoNumberSum(array []int, target int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] + array[j] == target {
				return []int{array[i], array[j]}
			}
		}
	}

	return []int{}
}