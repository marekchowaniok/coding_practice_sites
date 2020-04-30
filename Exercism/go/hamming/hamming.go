package hamming

import (
	"errors"
)

// Distance function which calculates hamming test
func Distance(a, b string) (int, error) {

	var count int
	if len(a) != len(b) {
		return 0, errors.New("length mismatch")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			count++
		}
	}

	return count, nil
}
