package GameOfLife

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Ground struct {
	Height, Width int
}

func (myGround *Ground) InitializeGround() [][]string {
	ground := make([][]string, myGround.Height)
	for i := 0; i < myGround.Height; i++ {
		ground[i] = make([]string, myGround.Width)
		for j := 0; j < myGround.Width; j++ {
			if rand.Intn(10)%2 == 0 {
				ground[i][j] = "*"
			} else {
				ground[i][j] = "-"
			}
		}
	}
	return ground
}

func PrintGround(ground [][]string) {
	fmt.Print("\n_______________________")
	for i := 0; i < len(ground); i++ {
		fmt.Print("\n")
		for j := 0; j < len(ground[i]); j++ {
			fmt.Printf("%v ", ground[i][j])
		}
	}
}

func GenerateNextGround(ground [][]string) [][]string {
	height := len(ground)
	nextGenGround := make([][]string, height)

	var totalAlive int
	var totalDead int

	for i := 0; i < height; i++ {
		width := len(ground[i])
		nextGenGround[i] = make([]string, width)
		for j := 0; j < width; j++ {
			totalAlive, totalDead = calculateTotalStats(ground, height, width)
			nextGenGround[i][j] = willBeAlive(CalculateAliveNeighbours(ground, i, j))
		}
	}
	printStats(totalAlive, totalDead)

	ground = nextGenGround
	return ground
}

func CalculateAliveNeighbours(ground [][]string, i, j int) (int, int) {
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

func calculateTotalStats(ground [][]string, Height, Width int) (totalAlive, totalDead int) {
	for i := 0; i < Height; i++ {
		for j := 0; j < Width; j++ {
			if ground[i][j] == "*" {
				totalAlive++
			} else {
				totalDead++
			}
		}
	}
	return totalAlive, totalDead
}

func printStats(totalAlive, totalDead int) {
	fmt.Printf("\n")
	fmt.Println("Total Alive Cells = ", totalAlive)
	fmt.Println("Total Dead Cells = ", totalDead)
}

func DelayBetweenGenerations(mytime int) {
	time.Sleep(time.Duration(mytime) * time.Second)
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
