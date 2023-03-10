package main

import "fmt"

func main() {
	var chars [7]int
	chars = [7]int{1071, 1040, 1064, 1040, 1056, 1042, 1054}
	for i := 0; i < 5; i++ {
		fmt.Printf("Nilai i = %d\n", i)
	}
	for j := 0; j < 11; j++ {
		if j == 5 {
			for k := 0; k < 7; k++ {
				fmt.Printf("Character %U '%c' starts at byte position %d\n", rune(chars[k]), chars[k], k*2)
			}
		} else {
			fmt.Printf("Nilai j = %d\n", j)
		}
	}
}
