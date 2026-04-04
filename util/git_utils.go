package util

import (
	"log"
	"os/exec"
	"strings"
	"os"
	"path"
	"fmt"
)

func GetRepoRoot() string {
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

	hookPath := path.Join(GetRepoRoot(), "/.git/hooks/", "pre-commit.sample")
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