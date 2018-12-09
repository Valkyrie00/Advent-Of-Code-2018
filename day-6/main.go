package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type coord struct {
	// id int
	X float64
	Y float64
}

type coords []coord

func main() {

	file := openFile("puzzle.txt")
	coords, maxWith, maxHeight := scanFile(file)

	inf := make(map[coord]bool)
	m := make(map[coord]int)
	// Per la parte 2
	distanceLessThan := float64(10000) // valore di partenza, deve essere minore di questo.
	nRegions := 0

	// Mi estendo fino all'estensione massima
	for y := float64(0); y < maxHeight; y++ {
		for x := float64(0); x < maxWith; x++ {
			mCoord := coord{0, 0}
			minDistance := float64(-1)
			totDistance := float64(0)

			for _, co := range coords {
				// Manhattan distance
				dist := math.Abs(x-co.X) + math.Abs(y-co.Y)
				if dist < minDistance || minDistance == -1 {
					minDistance = dist
					mCoord = co
				} else if dist == minDistance {
					mCoord = coord{-1, -1}
				}

				totDistance += math.Abs(x-co.X) + math.Abs(y-co.Y)
				// log.Println(totDistance)
			}

			if x == 0 || y == 0 || x == maxWith || y == maxHeight {
				inf[mCoord] = true
			}
			m[mCoord]++

			if totDistance < distanceLessThan {
				nRegions++
			}
		}
	}

	maxArea := 0
	for k, v := range m {
		if _, found := inf[k]; v > maxArea && !found {
			maxArea = v
		}
	}

	// Part 1
	fmt.Println("************ PART 1 ************")
	fmt.Println(maxArea)
	fmt.Println("********************************")

	// Part 2
	fmt.Println("************ PART 2 ************")
	fmt.Println(nRegions)
	fmt.Println("********************************")
}

func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func scanFile(file *os.File) (coords, float64, float64) {
	var line string
	scanner := bufio.NewScanner(file)

	var id int
	var coords []coord
	var maxWith float64
	var maxHeight float64

	for scanner.Scan() {
		line = scanner.Text()

		if len(line) > 0 {
			newCord := coord{}
			// newCord.id = id

			fmt.Sscanf(line, "%v, %v", &newCord.X, &newCord.Y)

			if newCord.X > maxWith {
				maxWith = newCord.X
			}
			if newCord.Y > maxHeight {
				maxHeight = newCord.Y
			}

			coords = append(coords, newCord)
			id++
		}
	}

	return coords, maxWith, maxHeight
}
