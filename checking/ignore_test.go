package checking

import (
	"testing"
)

func TestContainsIgnoreInlineComment(t *testing.T){
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
			got, err := ContainsIgnoreInlineComment(line)
			if err != nil {
				t.Errorf("error %s", err.Error())
			}
			if got != true{
				t.Errorf("Inline comment wasn't detected: %s", line)
			}
		})
	}
}