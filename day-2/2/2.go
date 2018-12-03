package main

import "fmt"

func main() {

	//TODO: KO

	boxIDS := []string{
		"abcde", "fghij", "fguij", "axcye", "fgmij",
	}

	var similar [][]string

	for _, ids := range boxIDS {

		//Prendo la prima
		var thisSimilar []string

		// fmt.Println("test")

		// thisLetterCount := make(map[string]int)
		var thisLetterCount []string

		for _, letter := range ids {
			l := string(letter)
			thisLetterCount = append(thisLetterCount, l)
		}

		// fmt.Println(thisLetterCount)

		thisSimilar = foundAnalogue(thisLetterCount, boxIDS)
		if len(thisSimilar) > 1 {
			thisSimilar = append(thisSimilar, ids)
			similar = append(similar, thisSimilar)
		}

		// Cerco quello simile
		// similar = append(similar, )

	}

	fmt.Println("--------------------------------------")
	fmt.Println("--------------END-------------")
	fmt.Println(similar)
	fmt.Println("--------------END -------------")
}

func foundAnalogue(thisLetterCount []string, boxIDS []string) []string {

	var similar []string

	fmt.Println("--------------------------------------")
	fmt.Println(thisLetterCount)
	fmt.Println("--------------------------------------")

	for _, ids := range boxIDS {

		// fmt.Println("verifico ids: " + ids)

		nCharDifferent := 0
		// var ciclyLetterCount []string
		for i, letter := range ids {
			l := string(letter)

			if thisLetterCount[i] != l {
				nCharDifferent++
			}
		}

		// if ids == "fghij" {
		fmt.Println(ids, nCharDifferent)
		// }

		if nCharDifferent < 2 {
			// non lo stesso
			if nCharDifferent != 0 {
				similar = append(similar, ids)
			}
		}
	}

	fmt.Println("SIMILAR:")
	fmt.Println(similar)

	// fmt.Println(similar)
	return similar
}
