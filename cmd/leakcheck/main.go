package main

import (
	"bufio"
	"fmt"
	"log"
	"github.com/traceback-afk/leakcheck/checking"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"flag"
)

type SecretOccurrence struct {
	File string
	Line int
	Text string
}

func GetGitRoot() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	out, err := cmd.Output()
	if err != nil {
		log.Fatal("Cannot get git root:", err)
	}
	return strings.TrimSpace(string(out))
}

func InstallHook(){
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	hook := fmt.Sprintf(`#!/bin/sh
# Pre-commit hook for key leak detection
"%s" --scan-staged
`, exePath)

	hookPath := path.Join(GetGitRoot(), "/.git/hooks/", "pre-commit.sample")
	err = os.WriteFile(hookPath, []byte(hook), 0755)
	if err != nil {
		log.Fatal("Failed to install hook:", err)
	}
	err = os.Rename(hookPath, strings.Replace(hookPath, ".sample", "", 1))
	if err != nil {
		log.Fatal("Error renaming the pre-commit file.", err)
	}
	fmt.Println("Pre-commit hook installed!")
}

func GetStagedDiff() (string, error) {
	cmd := exec.Command("git", "diff", "--cached")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
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
				relPath := strings.TrimPrefix(parts[2], "a/")
				currentFile = filepath.Join(repoRoot, relPath)
			}
			newStartLine = 0
		} else if strings.HasPrefix(line, "@@") {
			fmt.Sscanf(line, "@@ -%*d,%*d +%d,%*d @@", &newStartLine)
			currentLine = newStartLine - 1
		} else if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			addedLine := line[1:]
			currentLine++
			if checking.IsSecret(addedLine) {
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
	repoRoot := GetGitRoot()

	diff, err := GetStagedDiff()
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

func main() {
	installHook := flag.Bool("install-hook", false, "Install pre-commit hook")
	scanStaged := flag.Bool("scan-staged", false, "Scan staged changes")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: leakcheck [options]\n\n")
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
		fmt.Fprintf(flag.CommandLine.Output(), "\nExamples:\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  leakcheck --install-hook    Install the pre-commit hook\n")
		fmt.Fprintf(flag.CommandLine.Output(), "  leakcheck --scan-staged     Scan only staged changes\n")
	}
	flag.Parse()

	if !*installHook && !*scanStaged {
		flag.Usage()
		os.Exit(0)
	}

	if *installHook {
		InstallHook()
		return
	}

	if *scanStaged {
		ScanStaged()
		return
	}
}