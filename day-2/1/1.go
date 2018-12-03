package main

import "fmt"

func main() {

	boxIds := []string{
		"asgwdcmbrkerohqoutfylvzpnx", "asgwjcmbrkejihqoutfylvipne", "asgwjcmbrkedihqoutvylizpnz", "azgsjcmbrkedihqouifylvzpnx", "asgwucmbrktddhqoutfylvzpnx", "asgwocmbrkedihqoutfylvzivx", "aqgwjcmbrkevihqvutfylvzpnx", "tsgljcmbrkedihqourfylvzpnx", "asgpjcmbrkedihqoutfnlvzsnx", "astwjcmbrktdihqrutfylvzpnx", "asgwjcmbrpedhhqoutfylvzynx", "xsgwjcmbrkedieqowtfylvzpnx", "asgwjcmbvkedihfoutnylvzpnx", "asgwjcmtrkedihqouafylvzcnx", "asgwjcmbrkedihqoutfylvxpvm", "usgwjcmbrkedihqortfyuvzpnx", "asgwjcmbrwedihqoutfylizpix", "asgrjcvbrkedixqoutfylvzpnx", "asgwjcmbrogdihqoutfelvzpnx", "aggwjcmbrkesihqoutoylvzpnx", "asgtjccbrkedihqoutfrlvzpnx", "asgcucmbrbedihqoutfylvzpnx", "esgwjcmbrkedihqsutfylvzcnx", "asgwjcmbrkedrhqoutfyobzpnx", "mngwjcbbrkedihqoutfylvzpnx", "asgwjcrbrkeoihqyutfylvzpnx", "apgwjcmbrkednhqogtfylvzpnx", "asgwjcwbrkedihqoutfylplpnx", "asgwjcmbrkfdihqoutfxlvzpyx", "aegwjcmbrkedihqoutfylbxpnx", "asgljcmbrkedixqoutaylvzpnx", "aigwjcmbrkedihqouhfylvzpex", "asgwjbmbrkedihqoutfylfzpnp", "asgwjcmzroedihqoutcylvzinx", "asgwjcwbrieuihqoutfylvzpnx", "aagwjcmbrkedjhqdutfylvzpnx", "ahgwjcmbrkedihqsutfylvzpfx", "asgwjcmbrkedihzosttylvzpnx", "aegwjcmbrkedioqnutfylvzpnx", "asgwjcmbykidihqoutfysvzpnx", "asgwkcxbrkeddhqoutfylvzpnx", "ashwjcmbrkeeihqoutfylvzknx", "acgwjcmbrqedihqoqtfylvzpnx", "asgwjcmtrkedihooutfylszpnx", "asgwjcmbrkmdihqfutrylvzpnx", "asgwjcmbrkedihqoutjylvapnn", "asgwjcmbwkedihqoutkylkzpnx", "asgwjrmbrkedihqoutfycnzpnx", "asgwtcmbrkedihqoqtfylozpnx", "asgajcmbrkedihqoutuylvzpny", "asgwjcmbykedihqoutfylfzpwx", "asgwjcsbrkedihpoutfylvvpnx", "hsdwjcmbrvedihqoutfylvzpnx", "asgwjcmbrkedihqoutfdmszpnx", "adgwjcmbrtidihqoutfylvzpnx", "augwjcmbriedihqoutgylvzpnx", "asgwjvmbreedihqoutfllvzpnx", "asgwjcnbfkedihqoltfylvzpnx", "asgwjcmbykddihqoutqylvzpnx", "ajgwjcmbrkedihqoutfylvpvnx", "asgwjcmbrkydihqoutfylszpnl", "xsgwjcmbrkqdihqoutfylvkpnx", "asgwjcmbrkedimqoutfklvzknx", "csgwjbmbrkedihqoftfylvzpnx", "asgwjcmbjkedihjoutfylvzpnn", "asgwjcmprkedihqoulfalvzpnx", "asgwjcmbrvediqqoutfyuvzpnx", "asgwjambrkedhhqoutkylvzpnx", "asgejcmbrkidihqoutfylvzpnk", "hsiwjcmbrkedihqoutfylvzpnq", "asswjczbrkedihqoutfylczpnx", "asgwjnmbrkedyhzoutfylvzpnx", "asgwscmbrkedihqoutfklvlpnx", "asgwlcmbrktdihqoutfylvzpax", "asfwjcmerkedihqoutfylvipnx", "asgwjcmbrkeditqoeafylvzpnx", "asgwgcmbrkesihqoutfylyzpnx", "fsgwjcmbrkedihqouvfyavzpnx", "asgwjcmbrpedwhqoutfylmzpnx", "asgwjcmbrkzdzhqoucfylvzpnx", "asgwjcmnrketmhqoutfylvzpnx", "asgwjcmbrkedihxoutsylvzpnh", "asgwjcobrkedihqoutfrlvzpox", "asgwjcmbrkedihqootfylxzpox", "asgjjcmcrkedihqoutfylmzpnx", "lsgwjcmbrkedihqoutfyqvzunx", "asgwjcmbrwedihqoutoylvzpnu", "aszwjcmbtkedihqoutfylczpnx", "asgwjcmbykedihqoutfylvgpex", "asgijcmbrkedilqoutkylvzpnx", "astwxcmzrkedihqoutfylvzpnx", "akgwjcmbnkedihqfutfylvzpnx", "asgwjcmbrqndivqoutfylvzpnx", "asgwjrmbrleqihqoutfylvzpnx", "asgwjcmbrkevihqoutfxlvzpvx", "asbwjcmbrkedihqoutfelvwpnx", "asewjcbbrkmdihqoutfylvzpnx", "asgwjcmbrkeaihxoutfylpzpnx", "asgwjzmbrkedihqrotfylvzpnx", "asgwjcmbrkedihqoutgdxvzpnx", "asgwjcwbrkmdihqoutfylvzlnx", "asgwjcmbrkegihqoutfylrzpax", "ajgwjcmbrkegihqhutfylvzpnx", "asgwjcmbrzedihqhutfylvkpnx", "asgwjcmwrkedihqouhfylkzpnx", "aygwjcmbrkedihqoutfdlvzpnr", "asgwjcmbrkednhqoutiylvypnx", "aqgwjcmbrkezihqoutfylvzonx", "bsgwjcmbrkedihqouhfylvzsnx", "asgwjcmcrkedihqokyfylvzpnx", "asgsjcmbrkewiyqoutfylvzpnx", "asgwpcmbrkejihqoutfylzzpnx", "asgwjumbrkedbeqoutfylvzpnx", "asgwjcmbrkedihpoutqylqzpnx", "awgwjcmbrredihqoutfylvzpna", "asgwjsmbraedihqoutfylvzpvx", "asgwncmbrkedihqoutfyljzrnx", "asgwncmbrkedihqohtfylvzonx", "asgwjcmbrkedihqlutfylvypux", "asgwjcmbbkedihooutfylkzpnx", "asghjcmsryedihqoutfylvzpnx", "asgwjcmbrkevihqoulfzlvzpnx", "asggjcmbrkedizqoutfylvzknx", "asbwjcmbriedihqoutfylvmpnx", "asgwjcmbrkedqbqoutfylvzenx", "asgwjcmprkedihqoutfylvzknp", "asgwjcmbrkerihqoutfwlvzpno", "asgwjcmvrkesihqoutrylvzpnx", "asgzjcmbrkedihqoutfnlvbpnx", "asfwjcmbrkhdihqoutfylpzpnx", "asgwjcmbskedihqdutfyyvzpnx", "asgwjcmzrkedihqoutcylvzinx", "asgwjcmbrkedibqoutfylvjonx", "asgwjcmbrbedihqoutfylmzbnx", "asgwjcmbrkedhhqoutmylczpnx", "asgwjcmbrkbgihqoutzylvzpnx", "asgwjcfbrkedihqoupfyxvzpnx", "asiwjcmbzkedihqoutfyluzpnx", "asvwjcmbrkedihqoitfylvzpns", "asgwjcmxikedihqoutfyevzpnx", "asgwjcmbrkedioqoutfylvzwox", "asgwjcmbrkedivqoutjyuvzpnx", "asgwjcmbkkydihqrutfylvzpnx", "asgwjcmbrkxdiuqoutfylvopnx", "asgwjcmbrkedihqouthylvzpra", "asgwjcmbrzedimloutfylvzpnx", "asgwjcmbrkedmhqoulfytvzpnx", "asgwjcmbrkzdihqrutfysvzpnx", "ssgwjcmxrkedihqoutftlvzpnx", "asgwjcmbrkedihqoutfajvzynx", "asgwjcmbrkqdihqxuufylvzpnx", "asmwjcabrkedihqouxfylvzpnx", "asgwjcmbrkeeihqoatfycvzpnx", "asgwjcjbrgedjhqoutfylvzpnx", "asgljcmtrkedihqoutoylvzpnx", "asgwjcmbrkedigqouzfylvzpvx", "ajgvjcmbkkedihqoutfylvzpnx", "asgwjcmbrkedihqtugfygvzpnx", "asgbjcmbrkedihboftfylvzpnx", "asgwjwmbrkedihqontfylhzpnx", "asgwjfmhrkedihqoutfylvqpnx", "asgwjxmbrkedihqoutzylvzpnj", "asgwjcrlrkedihqoutfylvzpsx", "aygwjcmbrkedihqoutsylvzdnx", "zsgwjcmbrkedihjogtfylvzpnx", "asgwjxmbrkegihqoutfylvopnx", "asgwjcmbrkedihqhutfylvzcnr", "asgwicmbrkewihvoutfylvzpnx", "asqwjcmbvkedihqoutfylvzknx", "asgwjcmbrkedihqoktfyevzpnu", "asgwjcmbrkudihqoutfylqzznx", "asgwjdmbrkedihqoutfylvvdnx", "asgwjcmbrkwmihqautfylvzpnx", "asgwjcmbrxedihqoctfyldzpnx", "asgwjdmbrkedjhqoutfyfvzpnx", "asgwjcmtrzedihqoutfylvzpnm", "bpgwjcmbrmedihqoutfylvzpnx", "asgwjctbrkedihqoqtfynvzpnx", "askhjcmbrkedihqoutfylvzrnx", "asgkjcmblkehihqoutfylvzpnx", "asgwjjmbrkedvhqoutfhlvzpnx", "asgwjcmbrkedihqoupzylvzknx", "asgwjcmbukedchqoutfylizpnx", "askwjcmdrkedihqoutwylvzpnx", "asgwjcmbtkcdihloutfylvzpnx", "asgwjcmbrkedwgqoutvylvzpnx", "asmwjcmbrkedihqoutfylozpnc", "asgwjcmbriedibqouofylvzpnx", "asgnjcmcrkedihqoupfylvzpnx", "asgzjcmbrksdihqoutiylvzpnx", "asgwjcrbkkedihqouafylvzpnx", "asgwjcmbkkvdihqqutfylvzpnx", "astwjcqbrkedihqoutfylvzpvx", "asgwjcmhrkehihqoutfylvzvnx", "asgwjcmbraeduhqoutmylvzpnx", "asgwjcmbrkedihquutnylvzptx", "asgwjcmbrkedilqoftfylvzpnz", "akgwjmmbrkedihqoutfylxzpnx", "asgwjcmbrkhxikqoutfylvzpnx", "asgcjcmetkedihqoutfylvzpnx", "fsgwjcmsrkedihooutfylvzpnx", "gsgwjcmbrkedihdoutfylvzdnx", "asgwtccbrkedihqoutfylvwpnx", "asuwjcmbrkedihqcutfylvzpox", "asgwacmbrkodihqlutfylvzpnx", "asgwjcmbrkediuqoutfylqhpnx", "asgwjcmbrkwdrhqoutfylvzpno", "angwjcsblkedihqoutfylvzpnx", "aigwjcmbyoedihqoutfylvzpnx", "adgwjcmbrkedihqtutfylyzpnx", "asgwjzmbrkeeihqputfylvzpnx", "asgwjcmbrkwdihqoutfylvzpwc", "asgpjcmbrkgdihqbutfylvzpnx", "osgwjmmbrkedijqoutfylvzpnx", "asgjjcmbrkkdihqoutfylvzynx", "asgwjcnerwedihqoutfylvzpnx", "azgwhcmbrkedicqoutfylvzpnx", "asnwjcmbrsedihqoutfylvzpnm", "hsgwjcmgrkedihqoutfilvzpnx", "asgwscmbrkjdihqoutfylvzpnm", "asgbjbmbrkedhhqoutfylvzpnx", "aswwjcmtrkedihqjutfylvzpnx", "asgwicmbrbedihqoutfylvzpnm", "asgwjcubrkedihqoutfylvbnnx", "asvwjcmbrkehihqoutfylhzpnx", "gsgwjcmbrkedihqoutsklvzpnx", "asgwjcubikedihqoitfylvzpnx", "asgwjpmbskedilqoutfylvzpnx", "aigwjcmbrxedihqoutyylvzpnx", "asgwjcpbrkedihxoutfynvzpnx", "asgwjcmbrkegihqoutfklvzcnx", "asgwjvubrkedjhqoutfylvzpnx", "asgwjcabrkedihqoutfyivzplx", "asgwjcxbrkedihqgutfylvepnx", "asgwlcmbrkedihqoutqylvwpnx", "asgwjhmbrkydihqhutfylvzpnx", "asgwjcmbrkedihqoutfylwzone", "asgwycmbrkadihqoutuylvzpnx", "asgwjcybrkedihqoftfylvzpne", "asgwjcmnrkedihrodtfylvzpnx", "asgwicmwrkedihqoutfyovzpnx", "aqgwjlmbrkedilqoutfylvzpnx", "asgwjcmzskwdihqoutfylvzpnx", "asgwjcmdrkebihqoutfylvjpnx", "asgwjcmbrkpdihqoutfylxzphx", "asgwjcmbrkedixqoutlylfzpnx", "asgwjcmbrkadihqoutfylvlpdx", "asgejcmqrkedyhqoutfylvzpnx", "asgwjcmvroedihpoutfylvzpnx", "asgwjcmxrkedihqoutfyivzpmx",
	}

	checksum := checkIds(boxIds)

	fmt.Println(checksum)
}

func checkIds(boxIds []string) int32 {
	var countTwo int32
	var countThree int32

	for _, ids := range boxIds {

		lettersCount := make(map[string]int)

		for _, letter := range ids {
			l := string(letter)
			_, exist := lettersCount[l]

			if exist {
				lettersCount[l]++
			} else {
				lettersCount[l] = 1
			}
		}

		haveTwo, haveThree := countTwoAndThree(lettersCount)

		if true == haveTwo {
			countTwo++
		}

		if true == haveThree {
			countThree++
		}

	}

	return countTwo * countThree
}

func countTwoAndThree(lettersCount map[string]int) (bool, bool) {
	var haveTwo bool
	var haveThree bool

	haveTwo = false
	haveThree = false

	for _, value := range lettersCount {
		if value == 2 {
			haveTwo = true
		} else if value == 3 {
			haveThree = true
		}
	}

	return haveTwo, haveThree
}
