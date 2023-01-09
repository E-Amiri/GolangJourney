package main

import (
	"fmt"
)

func main() {
	var t []string
	var f []string

	height := 20
	width := 10
	s := make([][]string, height)
	for i := 0; i < height; i++ {
		s[i] = make([]string, width)
		for j := 0; j < width; j++ {
			if j%2 == 0 {
				s[i][j] = "*"
				t = append(t, s[i][j])
			} else {
				s[i][j] = "-"
				f = append(f, s[i][j])
			}

		}
	}
	fmt.Println(s, "\n")
	fmt.Println("Total Alive cells =", len(t))
	fmt.Println("Total Dead cells =", len(f))
}
