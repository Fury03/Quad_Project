package main

import "github.com/01-edu/z01"

func main() {
	QuadE(5, 3)
	QuadE(5, 1)
	QuadE(1, 1)
	QuadE(1, 5)
}

func QuadE(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for row := 1; row <= y; row++ {
		for col := 1; col <= x; col++ {
			if row == 1 && col == 1 {
				z01.PrintRune('A')
			} else if row == 1 && col == x {
				z01.PrintRune('C')
			} else if row == y && col == 1 {
				z01.PrintRune('C')
			} else if row == y && col == x {
				z01.PrintRune('A')
			} else if (row == 1 && col > 1 && col < x) || (row == y && col > 1 && col < x) {
				z01.PrintRune('B')
			} else if (col == 1 && row > 1 && row < y) || (col == x && row > 1 && row < y) {
				z01.PrintRune('B')
			} else {
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune('\n')
	}
}
