package main

import (
	"github.com/ldebruijn/advent-of-code/2022/input"
	"log"
	"strings"
)

func one() {
	raw := input.Load("cmd/3/input.txt")

	priorities := 0

	for _, line := range raw {
		length := len(line)

		compartmentOne := line[:length/2]
		compartmentTwo := line[length/2:]

		for _, char := range compartmentOne {
			priority := 0
			if strings.Contains(compartmentTwo, string(char)) {
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
