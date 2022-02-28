package test

import "testing"

func TestNoRepeatMaxLength(t *testing.T) {
	givens := []struct {
		input string
		ans   int
	}{
		//Normal cases
		{"abc", 3},
		{"pwrfkekpw", 6},
		{"pwwkew", 3},

		//Edge cases
		{"bbbbbb", 1},
		{"a", 1},
		{"abcabcabcd", 4},
		{"aba  b", 3},
		{"", 0},

		//Chinese cases
		{"一二一", 2},
		{"会回复回复哈圣诞福利上课的房间", 12},
	}

	for _, given := range givens {
		if actual := NoRepeatMaxLength(given.input); actual != given.ans {
			t.Errorf("Expect NoRepeatMaxLength(%s) = %d, but actually got %d", given.input, given.ans, actual)
		}
	}
}

func BenchmarkNoRepeatMaxLength(b *testing.B) {
	input, ans := "会回复回复哈圣诞福利上课的房间a", 14

	for i := 0; i < 15; i++ {
		input += input
	}
	b.Logf("input length: %d", len(input))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if actual := NoRepeatMaxLength(input); actual != ans {
			b.Errorf("Expect NoRepeatMaxLength ans =is %d, but actually got %d", ans, actual)
		}
	}
	//13,604058ns 12,896088ns 5,959347ns
}
