package main

import (
	"fmt"
)

var height, width int

func main() {
	var t []bool
	var f []bool

	height := 20
	width := 10
	s := make([][]bool, height)
	for i := 0; i < height; i++ {
		s[i] = make([]bool, width)
		for j := 0; j < width; j++ {
			if j%2 == 0 {
				s[i][j] = true
				t = append(t, s[i][j])
			} else {
				s[i][j] = false
				f = append(f, s[i][j])
			}

		}
	}
	fmt.Println(s, "\n")
	fmt.Println("Total Alive cells = ", len(t))
	fmt.Println("Total Dead cells =", len(f))
}
