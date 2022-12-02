package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
)

func main() {
	input := loadInput2()

	sort.Sort(sort.Reverse(sort.IntSlice(input)))

	total := 0

	for i := range input[:3] {
		total += input[i]
	}

	log.Printf("Most calories %v", input[:3])
	log.Printf("top 3 %d", total)
}

func loadInput2() []int {
	var input []int

	path, err := filepath.Abs("2022/cmd/1/input.txt")

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Println("Error reading input", err)
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	calories := 0

	for scanner.Scan() {
		text := scanner.Text()

		if text == "" {
			input = append(input, calories)
			calories = 0
			continue
		}

		value, err := strconv.Atoi(text)
		if err != nil {
			log.Println("Error converting calories", err)
			continue
		}

		calories += value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
