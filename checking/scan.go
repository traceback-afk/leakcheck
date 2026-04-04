package checking

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/traceback-afk/leakcheck/util"
)

type SecretOccurrence struct {
	File string
	Line int
	Text string
}

func ScanDiff(diff string, repoRoot string) []SecretOccurrence {
	scanner := bufio.NewScanner(strings.NewReader(diff))
	var currentFile string
	var currentLine int
	var newStartLine int
	secrets := []SecretOccurrence{}

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "diff --git") {
			parts := strings.Split(line, " ")
			if len(parts) >= 4 {
				currentFile = strings.TrimPrefix(parts[2], "a/")
			}
			newStartLine = 0
		} else if strings.HasPrefix(line, "@@") {
			fmt.Sscanf(line, "@@ -%*d,%*d +%d,%*d @@", &newStartLine)
			currentLine = newStartLine - 1
		} else if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			addedLine := line[1:]
			currentLine++

			if ContainsIgnoreComment(line) || addedLine == "" {
				continue
			}

			if IsSecret(addedLine) {
				secrets = append(secrets, SecretOccurrence{
					File: currentFile,
					Line: currentLine,
					Text: addedLine,
				})
			}
		} else if !strings.HasPrefix(line, "-") {
			currentLine++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println("error scanning diff:", err)
	}
	return secrets
}

func ScanStaged() {
	repoRoot := util.GetRepoRoot()

	diff, err := util.GetStagedDiff()
	if err != nil {
		panic(err)
	}

	secrets := ScanDiff(diff, repoRoot)

	if len(secrets) > 0 {
		for _, s := range secrets {
			fmt.Printf("Secret detected: %s:%d -> %s\n", s.File, s.Line, s.Text)
		}
		os.Exit(1)
	} else {
		fmt.Println("No secrets detected.")
	}
}