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
	fmt.Print("\n_______________________")
	for i := 0; i < len(ground); i++ {
		fmt.Print("\n")
		for j := 0; j < len(ground[i]); j++ {
			fmt.Printf("%v ", ground[i][j])
		}
	}
}

func generateNextGround(ground [][]string) [][]string {
	height := len(ground)
	nextGenGround := make([][]string, height)
	for i := 0; i < height; i++ {
		width := len(ground[i])
		nextGenGround[i] = make([]string, width)
		for j := 0; j < width; j++ {
			//(ground[i][j])
			nextGenGround[i][j] = willBeAlive(calculateAliveNeighbours(ground, i, j))
		}
	}
	ground = nextGenGround
	return ground
}

func calculateAliveNeighbours(ground [][]string, i, j int) (int, int) {
	aliveNeighbours := 0
	isAlive := 0

	for di := i - 1; di < i+2; di++ {
		if di < 0 || di >= len(ground) {
			continue
		}
		for dj := j - 1; dj < j+2; dj++ {
			if dj < 0 || dj >= len(ground[i]) {
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
	return isAlive, aliveNeighbours
}

func willBeAlive(isAlive, aliveNeighbours int) (nextGenGround string) {
	if isAlive == 1 && 1 < aliveNeighbours && aliveNeighbours < 4 {
		nextGenGround = "*"
	} else if isAlive == 0 && aliveNeighbours == 3 {
		nextGenGround = "*"
	} else {
		nextGenGround = "-"
	}
	return nextGenGround
}

/*func calculateTotalStats(ground string) (int, int) {
	var totalAlive []string
	var totalDead []string

	if ground == "*" {
		totalAlive = append(totalAlive, "*")
	} else {
		totalDead = append(totalDead, "-")
	}
	return len(totalAlive), len(totalDead)
}*/

/*func printStats(totalAlive, totalDead int) {
	fmt.Println("Total Alive Cells = ", (totalAlive))
	fmt.Println("Total Dead Cells = ", (totalDead))
}*/

func delayBetweenGenerations(mytime int) {
	time.Sleep(time.Duration(mytime) * time.Second)
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
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
		//calculateTotalStats()
		//printStats(s)
		delayBetweenGenerations(1)
		clearScreen()
	}
}
