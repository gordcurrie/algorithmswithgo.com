package module01

import (
	"strconv"
	"strings"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
func DecToBase(dec, base int) string {
	output := strings.Builder{}
	more := true
	for more {
		remainder := dec % base
		output.WriteString(convertInt(remainder))

		dec = dec / base
		if dec == 0 {
			more = false
		}
	}

	return Reverse(output.String())
}

func convertInt(i int) string {
	switch i {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
		return strconv.Itoa(i)
	case 10:
		return "A"
	case 11:
		return "B"
	case 12:
		return "C"
	case 13:
		return "D"
	case 14:
		return "E"
	case 15:
		return "F"
	}

	return ""
}
