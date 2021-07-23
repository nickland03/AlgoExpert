package main

func FirstNonRepeatingCharacter(str string) int {
    // use hash table to track each letter with its frequency
    m := make(map[rune]int)

    for _, char := range str {
        m[char] += 1
    }

    for idx, char := range str {
        if m[char] == 1 {
            return idx
        }
    }

    return -1
}
