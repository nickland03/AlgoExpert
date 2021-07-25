package array_of_products

import (
    "testing"
)

func BenchmarkArrayOfProducts(b *testing.B) {
    for n := 0; n < b.N; n++ {
        arr := []int{6,1,6,14,1,7,4,324,123,645,9,6,1,5,2,6,4,3,1,24,6,7}
        ArrayOfProducts(arr)
    }
}

func BenchmarkArrayOfProducts2(b *testing.B) {
    for n := 0; n < b.N; n++ {
        arr := []int{6,1,6,14,1,7,4,324,123,645,9,6,1,5,2,6,4,3,1,24,6,7}
        ArrayOfProducts2(arr)
    }
}
