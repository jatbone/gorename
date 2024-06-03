package renamer

import (
	"fmt"
	"os"
	"strings"

	helpers "github.com/jatbone/gorename/pkg/helpers"
)

func Undo(defaultLogFile string) error {
	logFile := "gorename-log.txt"
	if defaultLogFile != "" {
		logFile = defaultLogFile
	}

	lines, err := helpers.ReadLines(logFile)
	if err != nil {
		return err
	}

	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.Split(line, "->")
		from := parts[0]
		to := parts[1]

		if from == "" || to == "" {
			continue
		}

		if _, err := os.Stat(to); err != nil {
			continue
		}

		err := os.Rename(to, from)
		if err != nil {
			return err
		}

	}

	errRemoveFile := os.Remove(logFile)
	if errRemoveFile != nil {
		fmt.Println("error remove logfile")
	}

	return nil

}
