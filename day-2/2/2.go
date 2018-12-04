package main

import (
	"fmt"
	"strings"
)

func main() {

	// boxIDS := []string{
	// 	"asgwdcmbrkerohqoutfylvzpnx", "asgwdcmbrkerohqoutfylvzpbx", "asgwjcmbrkedihqoutvyliz123",
	// }

	boxIDS := []string{
		"asgwdcmbrkerohqoutfylvzpnx", "asgwjcmbrkejihqoutfylvipne", "asgwjcmbrkedihqoutvylizpnz", "azgsjcmbrkedihqouifylvzpnx", "asgwucmbrktddhqoutfylvzpnx", "asgwocmbrkedihqoutfylvzivx", "aqgwjcmbrkevihqvutfylvzpnx", "tsgljcmbrkedihqourfylvzpnx", "asgpjcmbrkedihqoutfnlvzsnx", "astwjcmbrktdihqrutfylvzpnx", "asgwjcmbrpedhhqoutfylvzynx", "xsgwjcmbrkedieqowtfylvzpnx", "asgwjcmbvkedihfoutnylvzpnx", "asgwjcmtrkedihqouafylvzcnx", "asgwjcmbrkedihqoutfylvxpvm", "usgwjcmbrkedihqortfyuvzpnx", "asgwjcmbrwedihqoutfylizpix", "asgrjcvbrkedixqoutfylvzpnx", "asgwjcmbrogdihqoutfelvzpnx", "aggwjcmbrkesihqoutoylvzpnx", "asgtjccbrkedihqoutfrlvzpnx", "asgcucmbrbedihqoutfylvzpnx", "esgwjcmbrkedihqsutfylvzcnx", "asgwjcmbrkedrhqoutfyobzpnx", "mngwjcbbrkedihqoutfylvzpnx", "asgwjcrbrkeoihqyutfylvzpnx", "apgwjcmbrkednhqogtfylvzpnx", "asgwjcwbrkedihqoutfylplpnx", "asgwjcmbrkfdihqoutfxlvzpyx", "aegwjcmbrkedihqoutfylbxpnx", "asgljcmbrkedixqoutaylvzpnx", "aigwjcmbrkedihqouhfylvzpex", "asgwjbmbrkedihqoutfylfzpnp", "asgwjcmzroedihqoutcylvzinx", "asgwjcwbrieuihqoutfylvzpnx", "aagwjcmbrkedjhqdutfylvzpnx", "ahgwjcmbrkedihqsutfylvzpfx", "asgwjcmbrkedihzosttylvzpnx", "aegwjcmbrkedioqnutfylvzpnx", "asgwjcmbykidihqoutfysvzpnx", "asgwkcxbrkeddhqoutfylvzpnx", "ashwjcmbrkeeihqoutfylvzknx", "acgwjcmbrqedihqoqtfylvzpnx", "asgwjcmtrkedihooutfylszpnx", "asgwjcmbrkmdihqfutrylvzpnx", "asgwjcmbrkedihqoutjylvapnn", "asgwjcmbwkedihqoutkylkzpnx", "asgwjrmbrkedihqoutfycnzpnx", "asgwtcmbrkedihqoqtfylozpnx", "asgajcmbrkedihqoutuylvzpny", "asgwjcmbykedihqoutfylfzpwx", "asgwjcsbrkedihpoutfylvvpnx", "hsdwjcmbrvedihqoutfylvzpnx", "asgwjcmbrkedihqoutfdmszpnx", "adgwjcmbrtidihqoutfylvzpnx", "augwjcmbriedihqoutgylvzpnx", "asgwjvmbreedihqoutfllvzpnx", "asgwjcnbfkedihqoltfylvzpnx", "asgwjcmbykddihqoutqylvzpnx", "ajgwjcmbrkedihqoutfylvpvnx", "asgwjcmbrkydihqoutfylszpnl", "xsgwjcmbrkqdihqoutfylvkpnx", "asgwjcmbrkedimqoutfklvzknx", "csgwjbmbrkedihqoftfylvzpnx", "asgwjcmbjkedihjoutfylvzpnn", "asgwjcmprkedihqoulfalvzpnx", "asgwjcmbrvediqqoutfyuvzpnx", "asgwjambrkedhhqoutkylvzpnx", "asgejcmbrkidihqoutfylvzpnk", "hsiwjcmbrkedihqoutfylvzpnq", "asswjczbrkedihqoutfylczpnx", "asgwjnmbrkedyhzoutfylvzpnx", "asgwscmbrkedihqoutfklvlpnx", "asgwlcmbrktdihqoutfylvzpax", "asfwjcmerkedihqoutfylvipnx", "asgwjcmbrkeditqoeafylvzpnx", "asgwgcmbrkesihqoutfylyzpnx", "fsgwjcmbrkedihqouvfyavzpnx", "asgwjcmbrpedwhqoutfylmzpnx", "asgwjcmbrkzdzhqoucfylvzpnx", "asgwjcmnrketmhqoutfylvzpnx", "asgwjcmbrkedihxoutsylvzpnh", "asgwjcobrkedihqoutfrlvzpox", "asgwjcmbrkedihqootfylxzpox", "asgjjcmcrkedihqoutfylmzpnx", "lsgwjcmbrkedihqoutfyqvzunx", "asgwjcmbrwedihqoutoylvzpnu", "aszwjcmbtkedihqoutfylczpnx", "asgwjcmbykedihqoutfylvgpex", "asgijcmbrkedilqoutkylvzpnx", "astwxcmzrkedihqoutfylvzpnx", "akgwjcmbnkedihqfutfylvzpnx", "asgwjcmbrqndivqoutfylvzpnx", "asgwjrmbrleqihqoutfylvzpnx", "asgwjcmbrkevihqoutfxlvzpvx", "asbwjcmbrkedihqoutfelvwpnx", "asewjcbbrkmdihqoutfylvzpnx", "asgwjcmbrkeaihxoutfylpzpnx", "asgwjzmbrkedihqrotfylvzpnx", "asgwjcmbrkedihqoutgdxvzpnx", "asgwjcwbrkmdihqoutfylvzlnx", "asgwjcmbrkegihqoutfylrzpax", "ajgwjcmbrkegihqhutfylvzpnx", "asgwjcmbrzedihqhutfylvkpnx", "asgwjcmwrkedihqouhfylkzpnx", "aygwjcmbrkedihqoutfdlvzpnr", "asgwjcmbrkednhqoutiylvypnx", "aqgwjcmbrkezihqoutfylvzonx", "bsgwjcmbrkedihqouhfylvzsnx", "asgwjcmcrkedihqokyfylvzpnx", "asgsjcmbrkewiyqoutfylvzpnx", "asgwpcmbrkejihqoutfylzzpnx", "asgwjumbrkedbeqoutfylvzpnx", "asgwjcmbrkedihpoutqylqzpnx", "awgwjcmbrredihqoutfylvzpna", "asgwjsmbraedihqoutfylvzpvx", "asgwncmbrkedihqoutfyljzrnx", "asgwncmbrkedihqohtfylvzonx", "asgwjcmbrkedihqlutfylvypux", "asgwjcmbbkedihooutfylkzpnx", "asghjcmsryedihqoutfylvzpnx", "asgwjcmbrkevihqoulfzlvzpnx", "asggjcmbrkedizqoutfylvzknx", "asbwjcmbriedihqoutfylvmpnx", "asgwjcmbrkedqbqoutfylvzenx", "asgwjcmprkedihqoutfylvzknp", "asgwjcmbrkerihqoutfwlvzpno", "asgwjcmvrkesihqoutrylvzpnx", "asgzjcmbrkedihqoutfnlvbpnx", "asfwjcmbrkhdihqoutfylpzpnx", "asgwjcmbskedihqdutfyyvzpnx", "asgwjcmzrkedihqoutcylvzinx", "asgwjcmbrkedibqoutfylvjonx", "asgwjcmbrbedihqoutfylmzbnx", "asgwjcmbrkedhhqoutmylczpnx", "asgwjcmbrkbgihqoutzylvzpnx", "asgwjcfbrkedihqoupfyxvzpnx", "asiwjcmbzkedihqoutfyluzpnx", "asvwjcmbrkedihqoitfylvzpns", "asgwjcmxikedihqoutfyevzpnx", "asgwjcmbrkedioqoutfylvzwox", "asgwjcmbrkedivqoutjyuvzpnx", "asgwjcmbkkydihqrutfylvzpnx", "asgwjcmbrkxdiuqoutfylvopnx", "asgwjcmbrkedihqouthylvzpra", "asgwjcmbrzedimloutfylvzpnx", "asgwjcmbrkedmhqoulfytvzpnx", "asgwjcmbrkzdihqrutfysvzpnx", "ssgwjcmxrkedihqoutftlvzpnx", "asgwjcmbrkedihqoutfajvzynx", "asgwjcmbrkqdihqxuufylvzpnx", "asmwjcabrkedihqouxfylvzpnx", "asgwjcmbrkeeihqoatfycvzpnx", "asgwjcjbrgedjhqoutfylvzpnx", "asgljcmtrkedihqoutoylvzpnx", "asgwjcmbrkedigqouzfylvzpvx", "ajgvjcmbkkedihqoutfylvzpnx", "asgwjcmbrkedihqtugfygvzpnx", "asgbjcmbrkedihboftfylvzpnx", "asgwjwmbrkedihqontfylhzpnx", "asgwjfmhrkedihqoutfylvqpnx", "asgwjxmbrkedihqoutzylvzpnj", "asgwjcrlrkedihqoutfylvzpsx", "aygwjcmbrkedihqoutsylvzdnx", "zsgwjcmbrkedihjogtfylvzpnx", "asgwjxmbrkegihqoutfylvopnx", "asgwjcmbrkedihqhutfylvzcnr", "asgwicmbrkewihvoutfylvzpnx", "asqwjcmbvkedihqoutfylvzknx", "asgwjcmbrkedihqoktfyevzpnu", "asgwjcmbrkudihqoutfylqzznx", "asgwjdmbrkedihqoutfylvvdnx", "asgwjcmbrkwmihqautfylvzpnx", "asgwjcmbrxedihqoctfyldzpnx", "asgwjdmbrkedjhqoutfyfvzpnx", "asgwjcmtrzedihqoutfylvzpnm", "bpgwjcmbrmedihqoutfylvzpnx", "asgwjctbrkedihqoqtfynvzpnx", "askhjcmbrkedihqoutfylvzrnx", "asgkjcmblkehihqoutfylvzpnx", "asgwjjmbrkedvhqoutfhlvzpnx", "asgwjcmbrkedihqoupzylvzknx", "asgwjcmbukedchqoutfylizpnx", "askwjcmdrkedihqoutwylvzpnx", "asgwjcmbtkcdihloutfylvzpnx", "asgwjcmbrkedwgqoutvylvzpnx", "asmwjcmbrkedihqoutfylozpnc", "asgwjcmbriedibqouofylvzpnx", "asgnjcmcrkedihqoupfylvzpnx", "asgzjcmbrksdihqoutiylvzpnx", "asgwjcrbkkedihqouafylvzpnx", "asgwjcmbkkvdihqqutfylvzpnx", "astwjcqbrkedihqoutfylvzpvx", "asgwjcmhrkehihqoutfylvzvnx", "asgwjcmbraeduhqoutmylvzpnx", "asgwjcmbrkedihquutnylvzptx", "asgwjcmbrkedilqoftfylvzpnz", "akgwjmmbrkedihqoutfylxzpnx", "asgwjcmbrkhxikqoutfylvzpnx", "asgcjcmetkedihqoutfylvzpnx", "fsgwjcmsrkedihooutfylvzpnx", "gsgwjcmbrkedihdoutfylvzdnx", "asgwtccbrkedihqoutfylvwpnx", "asuwjcmbrkedihqcutfylvzpox", "asgwacmbrkodihqlutfylvzpnx", "asgwjcmbrkediuqoutfylqhpnx", "asgwjcmbrkwdrhqoutfylvzpno", "angwjcsblkedihqoutfylvzpnx", "aigwjcmbyoedihqoutfylvzpnx", "adgwjcmbrkedihqtutfylyzpnx", "asgwjzmbrkeeihqputfylvzpnx", "asgwjcmbrkwdihqoutfylvzpwc", "asgpjcmbrkgdihqbutfylvzpnx", "osgwjmmbrkedijqoutfylvzpnx", "asgjjcmbrkkdihqoutfylvzynx", "asgwjcnerwedihqoutfylvzpnx", "azgwhcmbrkedicqoutfylvzpnx", "asnwjcmbrsedihqoutfylvzpnm", "hsgwjcmgrkedihqoutfilvzpnx", "asgwscmbrkjdihqoutfylvzpnm", "asgbjbmbrkedhhqoutfylvzpnx", "aswwjcmtrkedihqjutfylvzpnx", "asgwicmbrbedihqoutfylvzpnm", "asgwjcubrkedihqoutfylvbnnx", "asvwjcmbrkehihqoutfylhzpnx", "gsgwjcmbrkedihqoutsklvzpnx", "asgwjcubikedihqoitfylvzpnx", "asgwjpmbskedilqoutfylvzpnx", "aigwjcmbrxedihqoutyylvzpnx", "asgwjcpbrkedihxoutfynvzpnx", "asgwjcmbrkegihqoutfklvzcnx", "asgwjvubrkedjhqoutfylvzpnx", "asgwjcabrkedihqoutfyivzplx", "asgwjcxbrkedihqgutfylvepnx", "asgwlcmbrkedihqoutqylvwpnx", "asgwjhmbrkydihqhutfylvzpnx", "asgwjcmbrkedihqoutfylwzone", "asgwycmbrkadihqoutuylvzpnx", "asgwjcybrkedihqoftfylvzpne", "asgwjcmnrkedihrodtfylvzpnx", "asgwicmwrkedihqoutfyovzpnx", "aqgwjlmbrkedilqoutfylvzpnx", "asgwjcmzskwdihqoutfylvzpnx", "asgwjcmdrkebihqoutfylvjpnx", "asgwjcmbrkpdihqoutfylxzphx", "asgwjcmbrkedixqoutlylfzpnx", "asgwjcmbrkadihqoutfylvlpdx", "asgejcmqrkedyhqoutfylvzpnx", "asgwjcmvroedihpoutfylvzpnx", "asgwjcmxrkedihqoutfyivzpmx",
	}

	var similar []string

	for _, ids := range boxIDS {

		var thisSimilar []string
		var thisLetterCount []string

		for _, letter := range ids {
			l := string(letter)
			thisLetterCount = append(thisLetterCount, l)
		}

		thisSimilar = searchSimilar(thisLetterCount, boxIDS)
		if len(thisSimilar) >= 1 {
			thisSimilar = append(thisSimilar, ids)
			for _, sim := range thisSimilar {

				if !inSlice(sim, similar) {
					similar = append(similar, sim)
				}
			}
		}
	}

	// Clear strings
	if len(similar) == 2 {

		removeThisChar := foundDifference(strings.Split(similar[0], ""), strings.Split(similar[1], ""))
		fmt.Println(similar[0], similar[1], removeThisChar[0])
		trimmed := strings.Replace(similar[0], removeThisChar[0], "", -1)

		fmt.Println("answer: " + trimmed)
	}
}

func searchSimilar(thisLetterCount []string, boxIDS []string) []string {

	var similar []string

	// fmt.Println("--------------------------------------")
	// fmt.Println(thisLetterCount)
	// fmt.Println("--------------------------------------")

	for _, ids := range boxIDS {

		nCharDifferent := 0
		for i, letter := range ids {
			l := string(letter)
			if thisLetterCount[i] != l {
				nCharDifferent++
			}
		}

		// fmt.Println(ids, nCharDifferent)

		if nCharDifferent != 0 {
			if nCharDifferent < 2 {
				similar = append(similar, ids)
			}
		}
	}

	// fmt.Println("SIMILAR:")
	// fmt.Println(similar)

	return similar
}

func foundDifference(slice1 []string, slice2 []string) []string {
	var diff []string

	for i := 0; i < 2; i++ {
		for _, s1 := range slice1 {
			found := false
			for _, s2 := range slice2 {
				if s1 == s2 {
					found = true
					break
				}
			}

			if !found {
				diff = append(diff, s1)
			}
		}

		if i == 0 {
			slice1, slice2 = slice2, slice1
		}
	}

	return diff
}

func inSlice(s string, sls []string) bool {
	for _, sl := range sls {
		if s == sl {
			return true
		}
	}

	return false
}
