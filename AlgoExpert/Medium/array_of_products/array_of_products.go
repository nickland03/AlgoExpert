package array_of_products

func ArrayOfProducts(array []int) []int {
    output := make([]int, len(array))

    for i := 0; i < len(array); i++ {
        product := 1
        for j := 0; j < len(array); j++ {
            if i != j {
                product *= array[j]
            }
        }

        output[i] = product
    }

    return output
}

func ArrayOfProducts2(array []int) []int {
    products := make([]int, len(array))
    for idx, _ := range products {
        products[idx] = 1
    }

    leftProduct := 1
    for i := 0; i < len(array); i++ {
        products[i] = leftProduct
        leftProduct *= array[i]
    }

    rightProduct := 1
    for i := len(array) - 1; i >= 0; i-- {
        products[i] *= rightProduct
        rightProduct *= array[i]
    }

    return products
}
