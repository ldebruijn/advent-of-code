package main

import (
	"github.com/ldebruijn/advent-of-code/2022/input"
	"log"
)

func one() {
	raw := input.Load("cmd/7/input.txt")

	root, dirs := buildTree(raw)

	totalSize := 0

	for _, dir := range dirs {
		size := calculateSize(dir)
		if size <= 100_000 {
			totalSize += size
		}
	}

	log.Println("total size", totalSize)
	log.Println("root", root)
}
