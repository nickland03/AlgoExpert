package main

import "strconv"

func RunLengthEncoding(str string) string {
    count := 1
    var output []byte

    for i := 1; i < len(str); i++ {
        prev := str[i-1]
        curr := str[i]

        if prev != curr || count == 9 {
            output = append(output, strconv.Itoa(count)[0], prev)
            count = 0
        }

        count++
    }

    output = append(output, strconv.Itoa(count)[0], str[len(str)-1])

    return string(output)
}
