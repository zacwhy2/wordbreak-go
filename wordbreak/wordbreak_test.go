package wordbreak

import "testing"

func TestWordbreak(t *testing.T) {
	tests := []struct {
		s        string
		wordDict []string
		want     []string
	}{
		{
			"catsanddog",
			[]string{"cat", "cats", "and", "sand", "dog"},
			[]string{"cats", "and", "dog"},
		},
		{
			"catasanddog",
			[]string{"cat", "cats", "and", "sand", "dog"},
			[]string{"cat", "not possible", "sand", "dog"},
		},
		{
			"catsandadog",
			[]string{"cat", "cats", "and", "sand", "dog"},
			[]string{"cats", "and", "not possible", "dog"},
		},

		{
			"pineapplepenapple",
			[]string{"apple", "pen", "applepen", "pine", "pineapple"},
			[]string{"pineapple", "pen", "apple"},
		},
	}
	for _, tt := range tests {
		if got := WordBreak(tt.s, tt.wordDict); !equal(got, tt.want) {
			t.Errorf("WordBreak(%s, %v) = %v, want %v",
				tt.s,
				tt.wordDict,
				got,
				tt.want,
			)
		}
	}
}

// https://yourbasic.org/golang/compare-slices/

// equal tells whether a and b contain the same elements.
// A nil argument is equivalent to an empty slice.
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
