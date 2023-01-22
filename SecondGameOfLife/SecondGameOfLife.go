package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	height := 0
	width := 0

	fmt.Println("Enter the Height of Your Grid, Please ")
	fmt.Scanln(&height)
	fmt.Println("Enter the Width of Your Grid, Please ")
	fmt.Scanln(&width)

	s := make([][]string, height)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < height; i++ {
		s[i] = make([]string, width)
		fmt.Println("\n")
		for j := 0; j < width; j++ {
			if rand.Intn(10)%2 == 0 {
				s[i][j] = "*"
				fmt.Printf("%v ", s[i][j])
			} else {
				s[i][j] = "-"
				fmt.Printf("%v ", s[i][j])
			}
		}
	}
	fmt.Println("\n_______________________")

	for {
		for i := 0; i < height; i++ {
			fmt.Println("\n")
			for j := 0; j < width; j++ {
				alive := 0
				thisAlive := 0

				for di := i - 1; di < i+2; di++ {
					if di < 0 || di >= height {
						continue
					}
					for dj := j - 1; dj < j+2; dj++ {
						if dj < 0 || dj >= width {
							continue
						} else if s[di][dj] == "*" {
							if di == i && dj == j {
								thisAlive++
								//fmt.Printf("$")
							} else {
								alive++
								//fmt.Printf("#")
							}
						}
						//fmt.Printf("{%v}", di)
						//fmt.Printf("[%v]", dj)
					}
				}
				fmt.Printf("%v", alive)
				fmt.Printf("%v", thisAlive)
				if thisAlive == 1 && 1 < alive && alive < 4 {
					s[i][j] = "*"
					fmt.Printf("%v ", s[i][j])
				} else if thisAlive == 0 && alive == 3 {
					s[i][j] = "*"
					fmt.Printf("%v ", s[i][j])
				} else {
					s[i][j] = "-"
					fmt.Printf("%v ", s[i][j])
				}
			}
		}
		time.Sleep(1 * time.Second)
		fmt.Println("\n________________________")
	}
}
