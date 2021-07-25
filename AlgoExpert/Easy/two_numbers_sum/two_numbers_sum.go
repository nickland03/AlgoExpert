package main

import "sort"

func TwoNumberSum(array []int, target int) []int {
    sort.Ints(array)
    left := 0
    right := len(array) - 1

    for left < right {
        currSum := array[left] + array[right]
        if currSum == target {
            return []int{array[left], array[right]}
        } else if currSum < target {
            left += 1
        } else {
            right -= 1
        }
    }

    return []int{}
}
