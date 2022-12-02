package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const (
	Rock    = "R"
	Paper   = "P"
	Scissor = "S"

	Win  = "win"
	Draw = "draw"
	Loss = "loss"
)

type Game struct {
	Opponent string
	Own      string
}

func (g *Game) Score() int {
	score := 0

	if g.Own == Rock {
		score += 1
	} else if g.Own == Paper {
		score += 2
	} else if g.Own == Scissor {
		score += 3
	}

	conclusion := g.Conclusion()
	switch conclusion {
	case Win:
		score += 6
		break
	case Draw:
		score += 3
		break
	default:
		break
	}

	return score
}

func (g *Game) Conclusion() string {
	if g.Opponent == g.Own {
		return Draw
	} else if g.Opponent == Rock && g.Own == Paper {
		return Win
	} else if g.Opponent == Paper && g.Own == Scissor {
		return Win
	} else if g.Opponent == Scissor && g.Own == Rock {
		return Win
	} else {
		return Loss
	}
}

func one() {
	input := load()

	totalScore := 0

	for _, game := range input {
		totalScore += game.Score()
	}

	log.Println("Total score", totalScore)
}

func load() []Game {
	var input []Game

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

		game := Game{
			Opponent: toRockPaperScissor(values[0]),
			Own:      toRockPaperScissor(values[1]),
		}

		input = append(input, game)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}

func toRockPaperScissor(value string) string {
	switch value {
	case "A", "X":
		return Rock
	case "B", "Y":
		return Paper
	case "C", "Z":
		return Scissor
	default:
		return "Ohoh"
	}
	return ""
}
