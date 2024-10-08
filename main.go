package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/ZiplEix/gitign/generator"
	"github.com/ZiplEix/gitign/params"
)

func usage() {
	fmt.Println("Usage: gitign [language extension...] --ignore [path/to/folder...]")
	fmt.Println("Options:")
	fmt.Println("  -h          Show this help message")
	fmt.Println("  -v          Show the version")
	fmt.Println("  --ignore    Comma-separated list of things to ignore")
	fmt.Println("                  If no extension provided, will ignore the comma separated list of folders and extension | Example: gitign --ignore node_modules,build,.go")
	fmt.Println("  --append    Append the generated rules to an existing .gitignore file")
	fmt.Println("  --optimize  Remove duplicate rules from the .gitignore file (if used with no arguments, will optimize the .gitignore file)")
}

func main() {
	help := flag.Bool("h", false, "Show help message")
	ignoreFlag := flag.String("ignore", "", "Comma-separated list of things to ignore")
	appendFlag := flag.Bool("append", false, "Append the generated rules to an existing .gitignore file")
	optimizeFlag := flag.Bool("optimize", true, "Optimize the .gitignore by removing duplicate rules")
	versionFlag := flag.Bool("v", false, "Show the version")

	flag.Parse()

	if *help {
		usage()
		return
	}

	params := params.Params{
		Ignore:   strings.Split(*ignoreFlag, ","),
		Append:   *appendFlag,
		Optimize: *optimizeFlag,
		Version:  *versionFlag,
	}

	args := flag.Args()

	if params.Version {
		version()
		return
	}

	if len(args) == 1 && params.Optimize {
		println("Optimizing .gitignore")
		generator.OptimizeGitignore()
	}

	if len(args) == 0 {
		println("Detecting languages")
		generator.DetecteLanguages(params)
	} else {
		println("Generating gitignore")
		generator.GenerateGitignoreFromExtensions(args, params)
	}
}
