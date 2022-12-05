package main

import (
	"github.com/ldebruijn/advent-of-code/2022/input"
	"log"
	"strconv"
	"strings"
)

func two() {
	raw := input.Load("cmd/5/input.txt")

	stacks := map[int]Stack{}
	stacks[1] = make(Stack, 0)
	stacks[2] = make(Stack, 0)
	stacks[3] = make(Stack, 0)
	stacks[4] = make(Stack, 0)
	stacks[5] = make(Stack, 0)
	stacks[6] = make(Stack, 0)
	stacks[7] = make(Stack, 0)
	stacks[8] = make(Stack, 0)
	stacks[9] = make(Stack, 0)

	for i, line := range raw {
		if i < 8 {
			stacks = initialStacking(stacks, line)
			continue
		}

		if len(strings.TrimSpace(line)) == 0 {
			for i, stack := range stacks {
				stacks[i] = stack.Reverse()
			}
			continue
		}

		if strings.HasPrefix(line, "move ") {
			segments := strings.Split(line, " ")
			if len(segments) < 6 {
				log.Println("ohoh", line)
				continue
			}
			amountToMove, _ := strconv.Atoi(segments[1])
			from, _ := strconv.Atoi(segments[3])
			to, _ := strconv.Atoi(segments[5])

			stack, result := stacks[from].Take(amountToMove)
			stacks[from] = stack

			stacks[to] = stacks[to].PushMany(result)
		}
	}

	result := ""

	for i := 1; i <= len(stacks); i++ {
		stack := stacks[i]
		val := stack[len(stack)-1:]
		str := strings.ReplaceAll(val[0], "[", "")
		str = strings.ReplaceAll(str, "]", "")

		result += str
	}

	log.Println("result", result)
}
