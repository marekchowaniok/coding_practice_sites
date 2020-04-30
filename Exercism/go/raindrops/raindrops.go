package raindrops

import "strconv"

// Convert function running raindrops
func Convert(i int) string {
	var result string
	if i%3 == 0 {
		result += "Pling"
	}
	if i%5 == 0 {
		result += "Plang"
	}
	if i%7 == 0 {
		result += "Plong"
	}
	if len(result) == 0 {
		result = strconv.Itoa(i)
	}

	return result
}
