package checking

import (
	"bufio"
	"bytes"
	"log"
	"github.com/traceback-afk/leakcheck/parsing"
	"os"
	"runtime"
	"strings"
	"path/filepath"
)

var secretKeywords []string

// read keywords from file
func init() {
    _, filename, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatal("Cannot get caller info")
	}
	dir := filepath.Dir(filename)

	filePath := filepath.Join(dir, "variable-names.txt")

	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(bytes.NewReader(file))
    for scanner.Scan() {
        secretKeywords = append(secretKeywords, strings.ToLower(scanner.Text()))
    }

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ContainsSecretKeyword(varName string) bool {
	varName = strings.ToLower(varName)

	for _, value := range secretKeywords {
		value := strings.ToLower(value)
		if strings.Contains(varName, value) {
			return true
		}
	}

	return false
}

func IsSecret(line string) bool {
	score := 0
	key, value := parsing.ParseLine(line)

	// check if keyword is nearby
	keyContainsKW := false
	if key != "" {
		keyContainsKW = ContainsSecretKeyword(key)
	}
	valueContainsKW := ContainsSecretKeyword(value)
	if keyContainsKW || valueContainsKW {
		score += 5
	}

	// check length
	if len(value) >= 20 {
		score += 1
	}

	if CalculateEntropy(value) > 4.5 {
		score += 3
	}

	matchesRules, ruleName := MatchesRules(value)
	if matchesRules{
		score += 5
		log.Printf("Value matched: %s", ruleName)
	}

	if score >= 5 {
		return true
	}
	return false

}

