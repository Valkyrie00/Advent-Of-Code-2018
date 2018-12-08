package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// type unit string

func main() {

	file := openFile("puzzle.txt")
	polymer := scanFile(file)
	// polymer := "dabAcCaCBAcCcaDA" // Quiz Example
	var units []string

	// Inserisco caratteri in slice
	for _, unit := range polymer {
		units = append(units, string(unit))
	}

	haveReactions := false

	for {
		for i := range units {

			// Ultimo carattere non lo ciclo
			if i == len(units)-1 {
				haveReactions = false
				break
			}

			currentChar := i
			nextChar := i + 1

			if strings.ToUpper(units[currentChar]) == strings.ToUpper(units[nextChar]) && units[currentChar] != units[nextChar] {
				units = units[:currentChar+copy(units[currentChar:], units[nextChar+1:])]
				haveReactions = true
				break
			}
		}

		if haveReactions == false {
			break
		}
	}

	// Part 1
	fmt.Println("************ PART 1 ************")
	fmt.Println(len(units))
	fmt.Println("********************************")

}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func scanFile(file *os.File) string {
	var line string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}

	return line
}
