package main

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

    return SortThreeNumbers(a,b,c)
}

func SortThreeNumbers(a, b, c int) []int {
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