package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	lines []string
)

// GuardTimes - GuardTimes
type GuardTimes struct {
	Guard int
	Times map[int]int
}

func (t *GuardTimes) serachMax() (int, int) {
	var index, value int
	for i, v := range t.Times {
		if v > value {
			value = v
			index = i
		}
	}

	return index, value
}

// GuardsTimes - GuardsTimes
type GuardsTimes []GuardTimes

func main() {

	file := openFile("puzzle.txt")
	// Registro tutte le linee
	scan := scanFile(file)
	if !scan {
		panic("Errore durante lo scan :( ")
	}

	guardsFalls := make(map[int]int)
	guardsTotalTime := GuardsTimes{}
	var currentGuard int

	// Ciclo linee
	for i, line := range lines {
		if strings.Contains(line, "Guard") {
			currentGuard = parseGuard(line)

			// Se non esiste la guard la inserisco nuova e inizilizzo anche la nuova struct
			if _, ok := guardsFalls[currentGuard]; !ok {
				guardsTotalTime = append(guardsTotalTime, GuardTimes{Guard: currentGuard, Times: map[int]int{}})
			}

		} else if strings.Contains(line, "falls") {
			start := parseTime(line)
			end := parseTime(lines[i+1])
			diff := end - start

			// Per vedere chi è che dorme di più
			guardsFalls[currentGuard] = guardsFalls[currentGuard] + diff

			// Per vedere quando si addormentano di più
			for _, vg := range guardsTotalTime {
				if vg.Guard == currentGuard {
					for i := start; i < end; i++ {
						vg.Times[i]++
					}
				}
			}
		}
	}

	// test := GuardTimes{}

	// Prendo quello che dorme di più
	kingSleeperGuard := 0
	kingSleeperGuardVal := 0
	for guard, value := range guardsFalls {
		if value > kingSleeperGuardVal {
			kingSleeperGuard = guard
			kingSleeperGuardVal = value
		}
		// log.Println(guard)
	}

	// sort.Ints(guardsSleep)
	log.Println("King")
	log.Println(kingSleeperGuard)
	log.Println("----------------------")
	log.Println(guardsFalls)
	log.Println("----------------------")
	kingSleeperTime := 0
	kingSleeperTimeValTime := 0
	for _, vg := range guardsTotalTime {
		if vg.Guard == kingSleeperGuard {
			for time, value := range vg.Times {
				if value > kingSleeperTimeValTime {
					kingSleeperTime = time
					kingSleeperTimeValTime = value
				}
			}
			// log.Println(vg.Times)
		}
	}
	log.Println("King Time")
	log.Println(kingSleeperTime)
	fmt.Println("************ PART 1 ************")
	fmt.Println(kingSleeperGuard * kingSleeperTime)
	fmt.Println("********************************")

	// Part 2
	// var kingMaxMinute, kingMinute int

	// var maxGuard [60]int
	// var maxMinute [60]int

	var guardMinuteCache [60]int // most minutes
	var guardIDCache [60]int     // corresponding ID
	for _, vg := range guardsTotalTime {
		minute, val := vg.serachMax()
		if val > guardMinuteCache[minute] {
			guardMinuteCache[minute] = val
			guardIDCache[minute] = vg.Guard
		}
		// fmt.Println(minute, val)
	}

	// minute = 0
	// max = 0
	// guardID = 0
	var kingMaxMinute, kingMinute, kingGuad int
	for index, minutes := range guardMinuteCache {
		if minutes > kingMaxMinute {
			kingMaxMinute = minutes
			kingMinute = index
			kingGuad = guardIDCache[index]
		}
	}

	fmt.Println(kingGuad * kingMinute)

	// log.Println(kingMaxMinute)
	// log.Println(kingGuardMinute * kingMinute)
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
