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

type Node struct {
	name     string
	children []*Node
	parent   *Node
	size     int
}

func buildTree(raw []string) (Node, []*Node) {
	root := Node{
		name:     "/",
		children: []*Node{},
		parent:   nil,
		size:     0,
	}
	currentNode := &root

	var dirs []*Node

	for _, line := range raw {
		// command
		if line == "$ cd /" {
			continue
		}
		if strings.HasPrefix(line, "$") {
			if strings.HasPrefix(line, "$ cd") {
				//currentNode = changeDir(line, currentNode)
				path := strings.TrimPrefix(line, "$ cd ")

				// ..
				if path == "/" {
					continue
				} else if path == ".." {
					if currentNode.parent != nil {
						currentNode = currentNode.parent
					}
					continue
				}

				for _, child := range currentNode.children {
					if child.name == path {
						currentNode = child
						break
					}
				}
			}
		} else if strings.HasPrefix(line, "dir ") {
			dir := strings.TrimPrefix(line, "dir ")

			node := Node{
				name:     dir,
				children: []*Node{},
				parent:   currentNode,
				size:     0,
			}

			dirs = append(dirs, &node)

			currentNode.children = append(currentNode.children, &node)
		} else {
			// file
			segments := strings.Split(line, " ")

			if len(segments) < 2 {
				log.Println("Ohoh")
				continue
			}

			size, _ := strconv.Atoi(segments[0])

			node := Node{
				name:     segments[1],
				children: nil,
				parent:   nil,
				size:     size,
			}

			currentNode.children = append(currentNode.children, &node)
		}
	}

	return root, dirs
}

func calculateSize(node *Node) (size int) {
	if node.size > 0 {
		return node.size
	}

	for _, child := range node.children {
		size += calculateSize(child)
	}
	return
}
