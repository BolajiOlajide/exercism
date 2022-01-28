package grains

import (
	"errors"
	"math"
)

const MaxSquareNumber = 64
const MinSquareNumber = 1

func Square(number int) (uint64, error) {
	if number > MaxSquareNumber || number < MinSquareNumber {
		return 0, errors.New("Invalid square number")
	}

	exponent := number - 1
	return uint64(math.Pow(2, float64(exponent))), nil
}

func Total() uint64 {
	var sum uint64
	for idx := 1; idx <= MaxSquareNumber; idx++ {
		square, err := Square(idx)
		if err != nil {
			panic(err)
		}
		sum += square
	}

	return sum
}
