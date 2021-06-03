package main

func IsValidSubsequence(array []int, sequence []int) bool {
	sequenceIndex := 0

	for i := 0; i < len(array); i++ {
		if sequenceIndex == len(sequence) {
			break
		}

		if sequence[sequenceIndex] == array[i] {
			sequenceIndex++
		}
	}

	return sequenceIndex == len(sequence)
}