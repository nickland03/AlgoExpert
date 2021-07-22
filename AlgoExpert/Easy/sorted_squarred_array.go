package main

import (
    "sort"
)

func SortedSquaredArray(array []int) []int {
    arrLength := len(array)
    squaredArray := make([]int, arrLength)

    for i := 0; i < arrLength ; i++ {
        squaredNum := array[i] * array[i]
        squaredArray[i] = squaredNum
    }

    sort.Ints(squaredArray)
    return squaredArray
}