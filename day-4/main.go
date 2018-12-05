package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

var (
	lines []string
)

// Appunti
// verificare solo condizioni tra le 00.00 e le 00.59

func main() {
	file := openFile("puzzle.txt")
	// Registro tutte le linee
	scan := scanFile(file)
	if !scan {
		panic("Errore durante lo scan :( ")
	}
	log.Println(lines)
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func scanFile(file *os.File) bool {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Error
	if err := scanner.Err(); err != nil {
		log.Fatal(err)

		return false
	}

	sort.Strings(lines)

	return true
}
