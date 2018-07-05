package problem005

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	tests := []string{
		"babad",
		"cbbd",
		"123454321",
	}

	results := []string{
		"bab",
		"bb",
		"123454321",
	}
	caseNum := 3
	for i := 0; i < caseNum; i++ {
		if ret := longestPalindrome(tests[i]); ret != results[i] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}
	}
}