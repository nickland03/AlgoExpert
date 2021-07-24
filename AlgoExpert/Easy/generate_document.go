package main

type FreqTable map[rune]int
func GenerateUniqueFrequencyTableFromString(str string) FreqTable {
    m := FreqTable{}

    for _, char := range str {
        m[char]++
    }

    return m
}

func GenerateDocument(characters string, document string) bool {
    cmap := GenerateUniqueFrequencyTableFromString(characters)
    dmap := GenerateUniqueFrequencyTableFromString(document)

    for _, char := range document {
        if dmap[char] > cmap[char] {
            return false
        }
    }

    return true
}