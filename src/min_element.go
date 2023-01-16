package main

import "fmt"

func main() {
	x := []int{48, 2, 96, 86, 3, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17, 1}
	zero := 0
	for i, less := range x {
		if i == 0 {
			zero = less
		} else {
			if less < zero {
				zero = less
			}
		}
	}
	fmt.Println("Min integer: ", zero)
}
