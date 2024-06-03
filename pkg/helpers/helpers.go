package renamer

import (
	"bufio"
	"fmt"
	"os"
)

func WriteLines(filename string, lines []string) error {
	f, errOpen := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if errOpen != nil {
		return errOpen
	}
	for _, v := range lines {
		_, errWrite := fmt.Fprintln(f, v)
		if errWrite != nil {
			fmt.Println(errWrite)
		}
	}
	errClose := f.Close()
	if errClose != nil {
		return errClose
	}
	return nil
}

func ReadLines(logFile string) ([]string, error) {
	lines := make([]string, 0)

	f, errOpen := os.Open(logFile)

	if errOpen != nil {
		return []string{}, errOpen
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	return lines, nil
}
