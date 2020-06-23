package grains

import (
	"errors"
	"math"
)

//Square function for grains problem
func Square(in int) (uint64, error) {
	if in <= 0 || in > 64 {
		return uint64(0), errors.New("out of bounds")
	}
	result := uint64(math.Exp2(float64(in - 1)))
	return result, nil
}

//Total function which calculates grains
func Total() interface{} {
	max := 64
	var sum uint64
	for i := 2; i <= max; i++ {
		sq, err := Square(i)
		if err != nil {
			return errors.New("out of bounds")
		}
		sum += sq
	}
	return sum + 1
}
