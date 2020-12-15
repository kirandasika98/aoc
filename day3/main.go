package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Element string

const (
	Tree = Element("#")
	Open = Element(".")
)

var (
	maze [][]Element
	slop = []int{3, 1}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := buildRow(scanner.Text())
		fmt.Println(row)
		maze = append(maze, row)
	}
}

func buildRow(line string) []Element {
	row := []Element{}
	for _, c := range line {
		switch c {
		case '.':
			row = append(row, Element("."))
		case '#':
			row = append(row, Element("#"))
		}
	}

	return row
}

func findTreeCount(line string, index int) int {
	fmt.Println(index)
	max := strings.Index(line, "\n")
	if max == -1 {
		max = len(line)
	}
	treeCount := 0
	if string(line[index%max]) == "#" {
		treeCount = 1
	}

	if strings.Index(line, "\n") == -1 {
		return treeCount
	}
	return treeCount + findTreeCount(line[max+1:], index+slope)
}
