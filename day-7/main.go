package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type graph map[string][]string

func main() {
	file := openFile("puzzle.txt")

	scan := scanFile(file)

	log.Println(scan)
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func scanFile(file *os.File) bool {
	var line string
	scanner := bufio.NewScanner(file)

	graph := make(graph)

	for scanner.Scan() {
		line = scanner.Text()

		if len(line) > 0 {
			var start, end string
			fmt.Sscanf(line, "Step %s must be finished before step %s can begin.", &start, &end)

			if _, ok := graph[start]; ok {
				graph[start] = append(graph[start], end)
			} else {
				graph[start] = []string{end}

			}
		}
	}

	log.Println(graph)

	return true

	// return coords, maxWith, maxHeight
}
