package plur

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Text returns the correct form of the word based on the number.
func Text[N constraints.Integer](number N, one, two, many string) string {
	if number%100 >= 11 && number%100 <= 14 {
		return many
	}

	switch number % 10 {
	case 1:
		return one
	case 2, 3, 4:
		return two
	default:
		return many
	}
}

// NumberText returns the correct form of the word based on the number with the number itself.
func NumberText[N constraints.Integer](number N, one, two, many string) string {
	return fmt.Sprintf("%d %s", number, Text(number, one, two, many))
}
