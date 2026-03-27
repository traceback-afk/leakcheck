package checking

import (
	"testing"
)

func TestMatchesRules(t *testing.T) {

	tests := []struct {
		value		string
		expected	bool
	}{
		{"sdbfksjfd", false},
		{"ghp_abcdefghijklmnopqrstuvwx1234567890", true},
	}

	for _, tc := range tests {
		got, _ := MatchesRules(tc.value)
		if got != tc.expected {
			t.Errorf("value: %s, got: %v, expected: %v",
			tc.value, got, tc.expected )
		}
	}
}

