package raindrops

import (
    "strconv"
	"strings"
    )

var rules = []struct {
	divisor int
	word    string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

func Convert(number int) string {

	var sb strings.Builder

	for _, n := range rules {
		if number%n.divisor == 0 {
			sb.WriteString(n.word)
		}
	}

	if sb.Len() == 0 {
		return strconv.Itoa(number)
	}
	return sb.String()
}
