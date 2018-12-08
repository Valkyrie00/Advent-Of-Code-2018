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

// Guard -
type Guard struct {
	Guard int
	Falls int
	Times map[int]int
}

func (g *Guard) serachMax() (int, int) {
	var index, value int
	for i, v := range g.Times {
		if v > value {
			value = v
			index = i
		}
	}

	return index, value
}

func (g *Guard) incrementTimes(start int, end int) {
	for i := start; i < end; i++ {
		g.Times[i]++
	}
}

func (g *Guard) incrementFalls(value int) {
	g.Falls = g.Falls + value
}

func guardIsSet(guard int, guardList []Guard) bool {
	for _, g := range guardList {
		if g.Guard == guard {
			return true
		}
	}
	return false
}

// Guards - Guards
type Guards []Guard

func main() {
	guardsTotalTime := Guards{}
	var currentGuard int

	file := openFile("puzzle.txt")
	scan := scanFile(file)
	if !scan {
		panic("Errore durante lo scan :( ")
	}

	// Ciclo linee
	for i, line := range lines {
		if strings.Contains(line, "Guard") {
			currentGuard = parseGuard(line)

			// Se non esiste la guard la inserisco nuova e inizilizzo anche la nuova struct
			if ok := guardIsSet(currentGuard, guardsTotalTime); !ok {
				guardsTotalTime = append(guardsTotalTime, Guard{Guard: currentGuard, Falls: 0, Times: map[int]int{}})
			}

		} else if strings.Contains(line, "falls") {
			start, end, diff := parseFalls(i, line)

			for i := range guardsTotalTime {
				if guardsTotalTime[i].Guard == currentGuard {
					guardsTotalTime[i].incrementFalls(diff)
					guardsTotalTime[i].incrementTimes(start, end)
				}
			}
		}
	}

	// Prendo quello che dorme di più
	var kingSleeperGuard, kingSleeperGuardVal int
	for _, guard := range guardsTotalTime {
		if guard.Falls > kingSleeperGuardVal {
			kingSleeperGuard = guard.Guard
			kingSleeperGuardVal = guard.Falls
		}
	}

	log.Println(kingSleeperGuard, kingSleeperGuardVal)

	// Prendo quanto ha dormito di più
	var kingSleeperTime, kingSleeperTimeValTime int
	for _, vg := range guardsTotalTime {
		if vg.Guard == kingSleeperGuard {
			for time, value := range vg.Times {
				if value > kingSleeperTimeValTime {
					kingSleeperTime = time
					kingSleeperTimeValTime = value
				}
			}
		}
	}

	log.Println(kingSleeperTime, kingSleeperTimeValTime)

	fmt.Println("************ PART 1 ************")
	fmt.Println(kingSleeperGuard * kingSleeperTime)
	fmt.Println("********************************")

	// Part 2
	var guardMinuteCache, guardIDCache [60]int // most minutes
	for _, vg := range guardsTotalTime {
		minute, val := vg.serachMax()
		if val > guardMinuteCache[minute] {
			guardMinuteCache[minute] = val
			guardIDCache[minute] = vg.Guard
		}
	}

	var kingMaxMinute, kingMinute, kingGuad int
	for index, minutes := range guardMinuteCache {
		if minutes > kingMaxMinute {
			kingMaxMinute = minutes
			kingMinute = index
			kingGuad = guardIDCache[index]
		}
	}

	fmt.Println("************ PART 2 ************")
	fmt.Println(kingGuad * kingMinute)
	fmt.Println("********************************")
}

func parseGuard(line string) int {
	id, _ := strconv.Atoi(strings.Split(strings.Split(line, "#")[1], " ")[0])

	return id
}

func parseTime(line string) int {
	intTime, _ := strconv.Atoi(line[15:17])

	return intTime
}

func parseFalls(indexLine int, line string) (int, int, int) {
	start := parseTime(line)
	end := parseTime(lines[indexLine+1])
	diff := end - start

	return start, end, diff
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
