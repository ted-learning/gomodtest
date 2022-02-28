package test

func NoRepeatMaxLength(s string) int {
	lastOccurred := make([]int, 0xffff)
	for i := range lastOccurred {
		lastOccurred[i] = -1
	}
	start := 0
	maxLength := 0
	for i, c := range []rune(s) {
		if lastI := lastOccurred[c]; lastI != -1 && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[c] = i
	}
	return maxLength
}
