package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Ground struct {
	height, width int
}

func (myGround *Ground) initializeGround() [][]string {
	rand.Seed(time.Now().UnixNano())

	ground := make([][]string, myGround.height)
	for i := 0; i < myGround.height; i++ {
		ground[i] = make([]string, myGround.width)
		for j := 0; j < myGround.width; j++ {
			if rand.Intn(10)%2 == 0 {
				ground[i][j] = "*"
			} else {
				ground[i][j] = "-"
			}
		}
	}
	return ground
}

func printGround(ground [][]string) {
	for i := 0; i < len(ground); i++ {
		fmt.Println("\n")
		for j := 0; j < len(ground[i]); j++ {
			fmt.Printf("%v ", ground[i][j])
		}
	}
}
func generateNextGround(ground [][]string) [][]string { // too long
	fmt.Println("\n_______________________")
	var totalAlive []string
	var totalDead []string

	height := len(ground)
	nextGenGround := make([][]string, height)
	for i := 0; i < height; i++ {
		width := len(ground[i])
		//fmt.Print("\n")
		nextGenGround[i] = make([]string, width)
		for j := 0; j < width; j++ {
			aliveNeighbours := 0
			isAlive := 0
			for di := i - 1; di < i+2; di++ {
				if di < 0 || di >= height {
					continue
				}
				for dj := j - 1; dj < j+2; dj++ {
					if dj < 0 || dj >= width {
						continue
					} else if ground[di][dj] == "*" {
						if di == i && dj == j {
							isAlive++
						} else {
							aliveNeighbours++
						}
					}
				}
			}
			if isAlive == 1 && 1 < aliveNeighbours && aliveNeighbours < 4 {
				nextGenGround[i][j] = "*"
				totalAlive = append(totalAlive, "*")
			} else if isAlive == 0 && aliveNeighbours == 3 {
				nextGenGround[i][j] = "*"
				totalAlive = append(totalAlive, "*")
			} else {
				nextGenGround[i][j] = "-"
				totalDead = append(totalDead, "-")
			}
		}
	}
	ground = nextGenGround

	fmt.Println("Total Alive Cells = ", len(totalAlive))
	fmt.Println("Total Dead Cells = ", len(totalDead))
	return ground
}
func main() {
	var height, width int
	fmt.Println("Enter the Height of Your Grid, Please ")
	fmt.Scanln(&height)
	fmt.Println("Enter the Width of Your Grid, Please ")
	fmt.Scanln(&width)

	myGrid := Ground{height, width}
	ground := myGrid.initializeGround()
	printGround(ground)
	for {
		ground = generateNextGround(ground)
		printGround(ground)

		time.Sleep(1 * time.Second)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
