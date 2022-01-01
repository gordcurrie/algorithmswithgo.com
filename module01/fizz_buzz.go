package module01

import (
	"fmt"
	"strconv"
	"strings"
)

// FizzBuzz will print out all of the numbers
// from 1 to N replacing any divisible by 3
// with "Fizz", and divisible by 5 with "Buzz",
// and any divisible by both with "Fizz Buzz".
//
// Note: The test for this is a little
// complicated so that you can just use the
// `fmt` package and print to standard out.
// I wouldn't normally recommend this, but did
// it here to make life easier for beginners.
func FizzBuzz(n int) {
	sb := strings.Builder{}
	for i := 1; i <= n; i++ {
		var v string

		switch {
		case i%3 == 0 && i%5 == 0:
			v = "Fizz Buzz"
		case i%3 == 0:
			v = "Fizz"
		case i%5 == 0:
			v = "Buzz"
		default:
			v = strconv.Itoa(i)
		}

		if i == n {
			sb.WriteString(v)
		} else {
			sb.WriteString(fmt.Sprintf("%s, ", v))
		}
	}

	fmt.Println(sb.String())
}
