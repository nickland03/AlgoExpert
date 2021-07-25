package main

// BubbleSort Benchmark result | This implementation is better and faster
// BenchmarkBubbleSort-8 6700788 174.3 ns/op
func BubbleSort(array []int) []int {
    for i := 0; i < len(array)-1; i++ {
        for j := i + 1; j < len(array); j++ {
            if array[i] > array[j] {
                array[i], array[j] = array[j], array[i]
            }
        }
    }

    return array
}

// BubbleSort2 Benchmark result
// BenchmarkBubbleSort2-8 4900632 237.3 ns/op
func BubbleSort2(array []int) []int {
    isSorted := false
    counter := 0

    for !isSorted {
        isSorted = true
        for i := 0; i < len(array)-1-counter; i++ {
            if array[i] > array[i+1] {
                array[i], array[i+1] = array[i+1], array[i]
                isSorted = false
            }
        }

        counter++
    }

    return array
}