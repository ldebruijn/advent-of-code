package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var movesToScores = map[string]int{
	"A X": 3 + 0,
	"A Y": 1 + 3,
	"A Z": 2 + 6,
	"B X": 1 + 0,
	"B Y": 2 + 3,
	"B Z": 3 + 6,
	"C X": 2 + 0,
	"C Y": 3 + 3,
	"C Z": 1 + 6,
}

type Input struct {
	Opponent          string
	DesiredConclusion string
}

func two() {
	input := load2()

	totalScore := 0

	for _, game := range input {
		totalScore += movesToScores[fmt.Sprintf("%s %s", game.Opponent, game.DesiredConclusion)]
	}

	log.Println("Total score", totalScore)
}

func load2() []Input {
	var input []Input

	path, err := filepath.Abs("input.txt")

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		log.Println("Error reading input", err)
	}

	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		text := scanner.Text()

		values := strings.Split(text, " ")

		if len(values) < 2 {
			log.Println("Unexpected input", values)
			continue
		}

		game := Input{
			Opponent:          values[0],
			DesiredConclusion: values[1],
		}

		input = append(input, game)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
