package main

import (
	"flag"
	"fmt"

	renamer "github.com/jatbone/gorename/pkg/renamer"
	undo "github.com/jatbone/gorename/pkg/undo"
)

func main() {

	pattern := flag.String("pattern", "", "Regex pattern to match filenames")
	replacement := flag.String("replacement", "", "Replacement format for filenames")
	root := flag.String("root", "", "Base path for the files")
	dryRun := flag.Bool("dryRun", false, "Dry run")
	recursive := flag.Bool("recursive", false, "Execute the renaming recursively through all subdirectories")
	runUndo := flag.Bool("undo", false, "Undo rename operations")
	undoLogFile := flag.String("undoLogFile", "", "Undo rename operations from custom log file")

	flag.Parse()

	if *runUndo {
		fmt.Printf("Running undo operation")
		defaultLogFile := *undoLogFile

		if defaultLogFile != "" {
			fmt.Printf("Using logfile: %s\n", defaultLogFile)
		}

		err := undo.Undo(defaultLogFile)

		if err != nil {
			fmt.Println(err)
		}
	} else {

		fmt.Printf("Pattern: %s\n", *pattern)
		fmt.Printf("Replacement: %s\n", *replacement)
		fmt.Printf("Root: %s\n", *root)
		fmt.Printf("DryRun: %v\n", *dryRun)
		fmt.Printf("Recursive: %v\n", *recursive)

		err := renamer.Rename(renamer.RenamerParams{
			Pattern:     *pattern,
			Replacement: *replacement,
			Root:        *root,
			DryRun:      *dryRun,
			Recursive:   *recursive,
		})

		if err != nil {
			fmt.Println(err)
		}

	}

}
