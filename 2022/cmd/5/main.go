package main

import (
	"math"
	"strings"
)

func main() {
	one()
	two()
}

type Stack []string

func (s Stack) Push(value string) Stack {
	return append(s, value)
}

func (s Stack) Pop() (Stack, string) {
	if len(s) == 0 {
		return nil, ""
	}

	length := len(s)
	return s[:length-1], s[length-1]
}

func (s Stack) Take(num int) (Stack, []string) {
	if len(s) < num {
		num = int(math.Min(float64(len(s)), float64(num)))
	}

	return s[:len(s)-num], s[len(s)-num:]
}

func (s Stack) PushMany(value []string) Stack {
	return append(s, value...)
}

func (s Stack) Reverse() Stack {
	var res []string
	for i := len(s) - 1; i >= 0; i-- {
		res = append(res, s[i])
	}
	return res
}

func initialStacking(stacks map[int]Stack, line string) map[int]Stack {
	for i := 0; i < len(line)-2; i++ {
		if line[i:i+1] == " " {
			continue
		}

		value := line[i : i+3]

		if len(strings.TrimSpace(value)) == 0 {
			i += 3
			continue
		}

		stack := int(math.Floor(float64(i/4))) + 1

		stacks[stack] = stacks[stack].Push(value)
		i += 2
	}

	return stacks
}
