package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
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

	var Alive []string
	var Dead []string

	for i := 0; i < height; i++ {
		s[i] = make([]string, width)
		fmt.Println("\n")
		for j := 0; j < width; j++ {
			if rand.Intn(10)%2 == 0 {
				s[i][j] = "*"
				Alive = append(Alive, "*")
				fmt.Printf("%v ", s[i][j])
			} else {
				s[i][j] = "-"
				Dead = append(Dead, "-")
				fmt.Printf("%v ", s[i][j])
			}
		}
	}
	fmt.Println("Total Alive Cells = ", len(Alive))
	fmt.Println("Total Dead Cells = ", len(Dead))
	fmt.Println("\n_______________________")

	for {
		var totalAlive []string
		var totalDead []string
		nextGen := make([][]string, height)
		for i := 0; i < height; i++ {
			nextGen[i] = make([]string, width)
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
							} else {
								alive++
							}
						}

					}
				}
				if thisAlive == 1 && 1 < alive && alive < 4 {
					nextGen[i][j] = "*"
					totalAlive = append(totalAlive, "*")
					fmt.Printf("%v ", nextGen[i][j])
				} else if thisAlive == 0 && alive == 3 {
					nextGen[i][j] = "*"
					totalAlive = append(totalAlive, "*")
					fmt.Printf("%v ", nextGen[i][j])
				} else {
					nextGen[i][j] = "-"
					totalDead = append(totalDead, "-")
					fmt.Printf("%v ", nextGen[i][j])
				}
			}
		}
		s = nextGen
		time.Sleep(1 * time.Second)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		fmt.Println("Total Alive Cells = ", len(totalAlive))
		fmt.Println("Total Dead Cells = ", len(totalDead))
	}
}
