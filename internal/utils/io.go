package utils

import (
	"bufio"
	"fmt"
	"os"
)

// openFile opens a file and returns a reader or an error if the file doesn't exist
func OpenFile(path string) (*bufio.Reader, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("No such file or directory: %s", path)
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %v", err)
	}

	return bufio.NewReader(file), nil
}

// isInputFromPipe checks if the input is coming from a pipe (stdin)
func IsInputFromPipe() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false
	}
	return (stat.Mode() & os.ModeCharDevice) == 0 // Means it's from pipe
}
