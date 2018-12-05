package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	lines []string
	// guards map[int]int
	// guardsTime map[int][]int
)

// Appunti
// verificare solo condizioni tra le 00.00 e le 00.59

func main() {

	// guardsTime = append(guards)

	file := openFile("puzzle.txt")
	// Registro tutte le linee
	scan := scanFile(file)
	if !scan {
		panic("Errore durante lo scan :( ")
	}

	// guardsSleep := make(map[int]int)
	guardsTotalTime := make(map[int][]int)

	guardsTotalTime[0] = []int{2}
	guardsTotalTime[1] = []int{1}
	// guardsTime[1][5] = 1

	log.Println(guardsTotalTimex)

	// Prendo ciclando per guardia
	for index := 0; index < len(guardsTotalTime[1]); index++ {

		// Ciclo minuti per recupeare valore
		guardsTotalTimex := guardsTotalTime[index]
		log.Println(guardsTotalTimex)
	}

	//OK da qui in giù

	// var currentGuard int

	// for i, line := range lines {
	// 	// log.Println(line)

	// 	// var tempWakesUp int

	// 	// log.Println(currentGuard, tempWakesUp)

	// 	if strings.Contains(line, "Guard") {
	// 		currentGuard = parseGuard(line)

	// 	} else if strings.Contains(line, "falls") {
	// 		start := parseTime(line)
	// 		end := parseTime(lines[i+1])
	// 		diff := end - start

	// 		guardsSleep[currentGuard] = guardsSleep[currentGuard] + diff

	// 		//

	// 		// guardsTotalTime[currentGuard] = append(guardsTotalTime[currentGuard], map[int]int{parseTime(line): 2})

	// 		// log.Println(diff)
	// 	}
	// }

	// log.Println(guardsSleep)
}

func parseGuard(line string) int {
	id, _ := strconv.Atoi(strings.Split(strings.Split(line, "#")[1], " ")[0])

	return id
}

func parseTime(line string) int {
	intTime, _ := strconv.Atoi(line[15:17])

	return intTime
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
