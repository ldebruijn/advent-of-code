package main

import (
	"github.com/ldebruijn/advent-of-code/2022/input"
	"log"
	"sort"
)

func two() {
	raw := input.Load("cmd/7/input.txt")

	root, dirs := buildTree(raw)

	var validDirs []int

	for _, dir := range dirs {
		size := calculateSize(dir)
		if size > 8381165 {
			validDirs = append(validDirs, size)
		}
	}

	sort.Ints(validDirs)

	log.Println("root", root)
	log.Println("size", validDirs[0])
}

//8381165
//8854759
