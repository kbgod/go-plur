package plur

import "fmt"

type Number interface {
	uint8 | uint16 | uint32 | uint64 | int8 | int16 | int | int32 | int64
}

// Text returns the correct form of the word based on the number.
func Text[N Number](number N, one, two, many string) string {
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
func NumberText[N Number](number N, one, two, many string) string {
	return fmt.Sprintf("%d %s", number, Text(number, one, two, many))
}
