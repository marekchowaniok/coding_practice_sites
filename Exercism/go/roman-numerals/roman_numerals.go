package romannumerals

import (
	"errors"
	"sort"
)

var romanKeyMap = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

// ToRomanNumeral function for converting arabic to roman
func ToRomanNumeral(number int) (string, error) {
	if number < 1 || number > 3000 {
		return "", errors.New("Error")
	}

	keys := make([]int, 0)
	for k, _ := range romanKeyMap {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	result := ""
	for _, k := range keys {
		// fmt.Println(k)
		for number >= k {
			result += romanKeyMap[k]
			number -= k
		}
	}
	return result, nil
}
