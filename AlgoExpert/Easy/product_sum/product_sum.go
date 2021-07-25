package main

type SpecialArray []interface{}

func ProductSum(array SpecialArray) int {
    return RecursiveProductSum(array, 1)
}

func RecursiveProductSum(array SpecialArray, multiplier int) int {
    sum := 0

    for _, el := range array {
        if arr, ok := el.(SpecialArray); ok {
            sum += RecursiveProductSum(arr, multiplier+1)
        } else if num, ok := el.(int); ok {
            sum += num
        }
    }

    return sum * multiplier
}
