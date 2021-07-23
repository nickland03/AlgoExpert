package main

import "sort"

func FindThreeLargestNumbers(array []int) []int {
    a := array[0]
    b := array[1]
    c := array[2]

    for i := 3; i < len(array); i++ {
        currNum := array[i]
        if a <= b && a <= c && a < currNum {
            a = currNum
        } else if b <= a && b <= c && b < currNum {
            b = currNum
        } else if c <= a && c <= b && c < currNum {
            c = currNum
        }
    }

    // custom method to reorder 3 numbers in an array
    // or use SortThreeNumbers2(a,b,c) which uses a a built-in method
    return SortThreeNumbers1(a, b, c)
}

func SortThreeNumbers1(a, b, c int) []int {
    s := make([]int, 3)

    if a <= b && a <= c {
        if b <= c {
            s = []int{a, b, c}
        } else {
            s = []int{a, c, b}
        }
    }

    if b <= a && b <= c {
        if a <= c {
            s = []int{b, a, c}
        } else {
            s = []int{b, c, a}
        }
    }

    if c <= a && c <= b {
        if a <= b {
            s = []int{c, a, b}
        } else {
            s = []int{c, b, a}
        }
    }

    return s
}

func SortThreeNumbers2(a, b, c int) []int {
    s := []int{a, b, c}
    sort.Ints(s)

    return s
}
