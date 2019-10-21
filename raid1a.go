package piscine

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x <= 0 || y <= 0 {
		return
	} else {
		for i := 0; i < y; i++ {
			for j := 0; j < x; j++ {
				if (i == 0 && (j == x-1 || j == 0)) || (i == y-1 && (j == x-1 || j == 0)) {
					z01.PrintRune(111)
				} else if (j == 0 && (i != 0 || i != y-1)) || (j == x-1 && (i != 0 || i != y-1)) {
					z01.PrintRune(124)
				} else if i != 0 && j != 0 && i != y-1 && j != x-1 {
					z01.PrintRune(32)
				} else {
					z01.PrintRune(45)
				}

			}
			z01.PrintRune(10)
		}
	}
}
