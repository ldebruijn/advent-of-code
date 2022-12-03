package input

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
)

func Load(path string) []string {
	var input []string

	path, err := filepath.Abs(path)

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Println("Error reading input", err)
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
