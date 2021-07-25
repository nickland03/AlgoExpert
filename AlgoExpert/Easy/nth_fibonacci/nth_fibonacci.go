package main

func GetNthFib(n int) int {
	if n == 0 || n == 1 { return 0 }

	lastTwo := []int{0,1}

	for i := 2; i < n; i++ {
		next := lastTwo[0] + lastTwo[1]
		lastTwo = []int{lastTwo[1], next}
	}

	return lastTwo[1]
}