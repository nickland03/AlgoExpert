package main

import "testing"

func BenchmarkMinimumWaitingTime1(b *testing.B) {
    for n := 0; n < b.N; n++ {
        arr := []int{6,1,6,14,1,7,4,324,123,645,9,6,1,5,2,6,4,3,1,24,6,7}
        MinimumWaitingTime1(arr)
    }
}

func BenchmarkMinimumWaitingTime2(b *testing.B) {
    for n := 0; n < b.N; n++ {
        arr := []int{6,1,6,14,1,7,4,324,123,645,9,6,1,5,2,6,4,3,1,24,6,7}
        MinimumWaitingTime2(arr)
    }
}
