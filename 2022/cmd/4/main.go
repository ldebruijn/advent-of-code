package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	one()
	two()
}

type Slice []int

func toSection(section string) Slice {
	split := strings.Split(section, "-")

	if len(split) < 2 {
		log.Println("ohoh", split)
		return []int{}
	}

	var output = make([]int, 2)
	output[0], _ = strconv.Atoi(split[0])
	output[1], _ = strconv.Atoi(split[1])

	return output
}
