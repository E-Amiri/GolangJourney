package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Grid struct {
	height, width int
}

func (myGrid *Grid) makeGrid() [][]string {
	rand.Seed(time.Now().UnixNano())

	s := make([][]string, myGrid.height)
	for i := 0; i < myGrid.height; i++ {
		s[i] = make([]string, myGrid.width)
		for j := 0; j < myGrid.width; j++ {
			if rand.Intn(10)%2 == 0 {
				s[i][j] = "*"
			} else {
				s[i][j] = "-"
			}
		}
	}
	return s
}

func printGrid(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Println("\n")
		for j := 0; j < len(s[i]); j++ {
			fmt.Printf("%v ", s[i][j])
		}
	}
}
func nextGrid(s [][]string) [][]string {
	fmt.Println("\n_______________________")
	var totalAlive []string
	var totalDead []string

	height := len(s)
	nextGen := make([][]string, height)
	for i := 0; i < height; i++ {
		width := len(s[i])
		//fmt.Print("\n")
		nextGen[i] = make([]string, width)
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
			} else if thisAlive == 0 && alive == 3 {
				nextGen[i][j] = "*"
				totalAlive = append(totalAlive, "*")
			} else {
				nextGen[i][j] = "-"
				totalDead = append(totalDead, "-")
			}
		}
	}
	s = nextGen

	fmt.Println("Total Alive Cells = ", len(totalAlive))
	fmt.Println("Total Dead Cells = ", len(totalDead))
	return s
}
func main() {
	var h, w int
	fmt.Println("Enter the Height of Your Grid, Please ")
	fmt.Scanln(&h)
	fmt.Println("Enter the Width of Your Grid, Please ")
	fmt.Scanln(&w)

	myGrid := Grid{h, w}
	s := myGrid.makeGrid()
	printGrid(s)
	for {
		s = nextGrid(s)
		printGrid(s)

		time.Sleep(1 * time.Second)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
