package main

func CaesarCipherEncryptor(str string, key int) string {
    output := []rune(str)

    for idx, char := range str {
        newLetterAscii := int(char) + key%26
        if newLetterAscii > int('z') {
            newLetterAscii = int('a') - 1 + newLetterAscii - int('z')
        }
        newLetter := rune(newLetterAscii)
        output[idx] = newLetter
    }

    return string(output)
}
