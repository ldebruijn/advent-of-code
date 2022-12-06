package main

import (
	"github.com/ldebruijn/advent-of-code/2022/input"
	"log"
	"math"
	"strings"
)

func two() {
	raw := input.Load("cmd/6/input.txt")
	if len(raw) != 1 {
		log.Println("Unexpected input")
		return
	}

	sequence := raw[0]
	var marker string
	var index int

	for i := 0; i < len(sequence); i++ {
		start := int(math.Max(0, float64(i-14)))
		end := int(math.Min(float64(i), float64(len(sequence))))

		if end-start < 14 {
			// not enough to be our potentialMarker
			continue
		}

		potentialMarker := sequence[start:end]
		hasDuplicate := false

		for i, char := range potentialMarker {
			s := string(char)
			if i < 14 && strings.Count(potentialMarker[i+1:], s) > 0 {
				hasDuplicate = true
				break
			}
		}

		if !hasDuplicate {
			marker = potentialMarker
			index = i
			break
		}
	}

	log.Println("Marker", marker, "index", index)
}
