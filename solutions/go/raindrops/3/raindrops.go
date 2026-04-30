package raindrops

import (
    "strconv"
	"strings"
    )

func Convert(number int) string {
	var sb strings.Builder

    type pairs struct {
        divisor int
        word string
    }
    
    rules:= []pairs{
        {3, "Pling"},
        {5, "Plang"},
        {7, "Plong"},
    }

    for _, n := range rules {
        if number % n.divisor == 0 {
            sb.WriteString(n.word)
        }
    }
    
    if sb.Len() == 0 {
        return strconv.Itoa(number)
    }
	return sb.String()
}