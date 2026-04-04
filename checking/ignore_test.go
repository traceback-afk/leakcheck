package checking

import (
	"testing"
)

func TestContainsIgnoreComment(t *testing.T){
	lines := []string {
		"somepassword // leakcheck:ignore",
		"testpassword // leakcheck: ignore",
		"testpassword // leakcheck :ignore",
		"testpassword // leakcheck : ignore",
		"testpassword //leakcheck:ignore",
		"testpassword //leakcheck :ignore",
		"testpassword //leakcheck: ignore",
		"testpassword //leakcheck : ignore",
		"testpassword//leakcheck : ignore ",
	}

	for _, line := range lines {
		t.Run(line, func(t *testing.T) {
			got := ContainsIgnoreComment(line)
			if got != true{
				t.Errorf("Inline comment wasn't detected: %s", line)
			}
		})
	}
}