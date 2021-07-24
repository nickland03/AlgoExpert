package main

import "sort"

func MinimumWaitingTime1(queries []int) int {
    sort.Ints(queries)
    minimumWaitingTime := 0

    for i := 1; i < len(queries); i++ {
        for j := 0; j < i; j++ {
            minimumWaitingTime += queries[j]
        }
    }

    return minimumWaitingTime
}

func MinimumWaitingTime2(queries []int) int {
    sort.Ints(queries)
    minimumWaitingTime := 0

    for i := 0; i < len(queries); i++ {
        queriesLeft := len(queries) - (i + 1)
        minimumWaitingTime += queries[i] * queriesLeft
    }

    return minimumWaitingTime
}
