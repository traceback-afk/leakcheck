package parsing

import (
	"strings"
)

func ParseLine(line string) (key string, value string){
	parts := strings.SplitN(line, "=", 2)
	
	if len(parts) != 2 {
		return "", value
	}

	key, value = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])

	trimChars := []string{"'", "\""}

	for _, char := range trimChars {
		if strings.Contains(value, char){
			value = strings.Trim(value, char)
		}
	}

	return key, value
}