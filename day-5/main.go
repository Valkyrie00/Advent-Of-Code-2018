package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	file := openFile("puzzle.txt")
	polymer := scanFile(file)
	// polymer := "dabAcCaCBAcCcaDA" // Quiz Example

	// Part 1
	fmt.Println("************ PART 1 ************")
	fmt.Println(len(reduce(polymer)))
	fmt.Println("********************************")

	alphabet := "abcdefghijklmnopqrstuvwxyz"
	var allNewPolymerLen []int
	for _, abc := range alphabet {
		// Rimuovo dal token tutti i caratteri maiuscoli e minuscoli
		newPolymer := strings.Replace(strings.Replace(polymer, string(strings.ToUpper(string(abc))), "", -1), string(abc), "", -1)
		newPolymerLen := len(reduce(newPolymer))

		allNewPolymerLen = append(allNewPolymerLen, newPolymerLen)
	}

	shortestPolymer := len(polymer)
	for _, singlePolymer := range allNewPolymerLen {
		if singlePolymer < shortestPolymer {
			shortestPolymer = singlePolymer
		}
	}

	// Part 2
	fmt.Println("************ PART 1 ************")
	fmt.Println(shortestPolymer)
	fmt.Println("********************************")
}

func reduce(polymer string) string {
	var units []string

	// Inserisco caratteri in slice
	for _, unit := range polymer {
		units = append(units, string(unit))
	}

	reducedUnits := units
	haveReactions := false
	for {
		for i := range reducedUnits {

			// Ultimo carattere non lo ciclo
			if i == len(reducedUnits)-1 {
				haveReactions = false
				break
			}

			currentChar := i
			nextChar := i + 1

			if strings.ToUpper(reducedUnits[currentChar]) == strings.ToUpper(reducedUnits[nextChar]) && reducedUnits[currentChar] != reducedUnits[nextChar] {
				reducedUnits = reducedUnits[:currentChar+copy(reducedUnits[currentChar:], reducedUnits[nextChar+1:])]
				haveReactions = true
				break
			}
		}

		if haveReactions == false {
			break
		}
	}

	result := ""
	for _, char := range reducedUnits {
		result += char
	}

	return result
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
