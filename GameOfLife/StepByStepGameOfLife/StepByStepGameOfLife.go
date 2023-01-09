package main

import (
	"fmt"
)

func main() {
	var t []string
	var f []string
	height := 20
	width := 20

	fmt.Println("Please Enter the Height of Your Grid:")
	fmt.Scan(&height)
	fmt.Println("Please Enter the Width of Your Grid:")
	fmt.Scan(&width)

	s := make([][]string, height)
	for i := 0; i < height; i++ {
		s[i] = make([]string, width)
		fmt.Printf("\n")
		for j := 0; j < width; j++ {
			if j%2 == 0 {
				s[i][j] = "*"
				t = append(t, s[i][j])
				fmt.Printf("%v", s[i][j])
			} else {
				s[i][j] = "-"
				f = append(f, s[i][j])
				fmt.Printf("%v", s[i][j])
			}

		}
	}
	fmt.Println("\nTotal Alive cells =", len(t))
	fmt.Println("Total Dead cells =", len(f))
}
