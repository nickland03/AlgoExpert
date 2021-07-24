package main

import "sort"

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
    SortArrayAsc(redShirtSpeeds)

    if fastest {
        SortArrayDesc(blueShirtSpeeds)
    } else {
        SortArrayAsc(blueShirtSpeeds)
    }

    total := 0

    for i := 0; i < len(redShirtSpeeds); i++ {
        total += max(redShirtSpeeds[i], blueShirtSpeeds[i])
    }

    return total
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}

func SortArrayAsc(arr []int) {
    sort.Slice(arr, func(i, j int) bool {
        return arr[j] < arr[i]
    })
}

func SortArrayDesc(arr []int) {
    sort.Slice(arr, func(i, j int) bool {
        return arr[j] > arr[i]
    })
}
