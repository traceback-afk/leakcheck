package main

import (
	"fmt"
	"os"
	"flag"
	"github.com/traceback-afk/leakcheck/checking"
	"github.com/traceback-afk/leakcheck/util"
)


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
		util.InstallHook()
		return
	}

	if *scanStaged {
		checking.ScanStaged()
		return
	}
}