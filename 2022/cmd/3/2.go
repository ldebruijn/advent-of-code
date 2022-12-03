package main

import (
	"github.com/ldebruijn/advent-of-code/2022/input"
	"log"
	"strings"
)

func two() {
	raw := input.Load("cmd/3/input.txt")

	priorities := 0

	for i := 0; i < len(raw); i += 3 {
		first := raw[i]
		second := raw[i+1]
		third := raw[i+2]

		for _, char := range first {
			value := string(char)
			if strings.Contains(second, value) && strings.Contains(third, value) {
				priority := 0
				if char > 96 {
					// lowercase unicode numeric value
					priority = int(char) - 96
				} else if char > 64 {
					// uppercase unicode numeric value
					priority = int(char) - 64 + 26
				} else {
					log.Println("ohoh")
				}
				//log.Println("Char", char, string(char), priority)
				priorities += priority
				break
			}
		}
	}
	log.Println("Priorities", priorities)
}
