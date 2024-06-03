package renamer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	helpers "github.com/jatbone/gorename/pkg/helpers"
)

type RenamerParams struct {
	Pattern     string
	Replacement string
	Root        string
	DryRun      bool
	Recursive   bool
	Log         bool
	LogFile     string
}

func Rename(params RenamerParams) error {
	if params.Pattern == "" {
		return errors.New("Pattern string is required")
	}
	if params.Replacement == "" {
		return errors.New("Replacement string is required")
	}
	pattern := params.Pattern
	replacement := params.Replacement
	root := "."
	dryRun := false
	recursive := false
	logFile := "gorename-log.txt"
	if params.Root != "" {
		root = params.Root
	}
	if params.DryRun {
		dryRun = params.DryRun
	}
	if params.Recursive {
		recursive = params.Recursive
	}

	if params.LogFile != "" {
		logFile = params.LogFile
	}

	r, compErr := regexp.Compile(pattern)
	if compErr != nil {
		return errors.New("Failed to compile pattern")
	}

	var renames []string

	walkErr := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if info.IsDir() {
			return nil
		}

		relativePath, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		if filepath.Dir(relativePath) != "." && !recursive {
			return nil
		}

		filename := info.Name()
		matchString := r.MatchString(filename)

		if matchString {
			dir := filepath.Dir(path)
			newPath := filepath.Join(dir, r.ReplaceAllString(filename, replacement))

			if !dryRun {
				err := os.Rename(path, newPath)
				if err != nil {
					return err
				}
				renames = append(renames, path+"->"+newPath)
			} else {
				fmt.Println(path + "->" + newPath)
			}
		}

		return nil
	})

	if walkErr != nil {
		return walkErr
	}

	if len(renames) > 0 {
		writeRenamesErr := helpers.WriteLines(logFile, renames)
		if writeRenamesErr != nil {
			return writeRenamesErr
		}
	}

	return nil
}
