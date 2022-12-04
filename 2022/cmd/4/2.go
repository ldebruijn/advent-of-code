package main

import (
	"github.com/ldebruijn/advent-of-code/2022/input"
	"log"
	"strings"
)

func two() {
	raw := input.Load("cmd/4/input.txt")

	numFullContained := 0

	for _, line := range raw {
		sections := strings.Split(line, ",")
		if len(sections) < 2 {
			log.Println("Ohoh", line)
			continue
		}

		firstSection := toSection(sections[0])
		secondSection := toSection(sections[1])

		if firstSection.partiallyContains(secondSection) || secondSection.partiallyContains(firstSection) {
			numFullContained++
		}
	}

	log.Println("Fully contains", numFullContained)

}

func (s Slice) partiallyContains(other []int) bool {
	if len(s) < 2 || len(other) < 2 {
		return false
	}
	return s[0] <= other[0] && other[0] <= s[1]
}
