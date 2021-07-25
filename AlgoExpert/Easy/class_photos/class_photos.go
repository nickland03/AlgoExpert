package main

import "sort"

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
    SortRows(redShirtHeights, blueShirtHeights)

    if redShirtHeights[0] == blueShirtHeights[0] {
        return false
    }

    isRedLeading := false
    if redShirtHeights[0] > blueShirtHeights[0] {
        isRedLeading = true
    }

    for i := 0; i < len(redShirtHeights); i++ {
        if isRedLeading {
            if redShirtHeights[i] < blueShirtHeights[i] {
                return false
            }
        } else {
            if blueShirtHeights[i] < redShirtHeights[i] {
                return false
            }
        }
    }

    return true
}

func SortRows(red, blue []int) {
    sort.Slice(red, func(i, j int) bool {
        return red[i] < red[j]
    })

    sort.Slice(blue, func(i, j int) bool {
        return blue[i] < blue[j]
    })
}
