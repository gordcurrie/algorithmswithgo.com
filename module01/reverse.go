package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	reversed := make([]rune, len(word))

	for i, r := range word {
		reversed[len(reversed)-i-1] = r
	}

	return string(reversed)
}
