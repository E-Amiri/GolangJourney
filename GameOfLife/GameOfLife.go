package GameOfLife

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

func (myGround *Ground) InitializeGround() [][]string {
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

	var totalAlive []string
	var totalDead []string

	printStats(totalAlive, totalDead)

	for i := 0; i < height; i++ {
		width := len(ground[i])
		nextGenGround[i] = make([]string, width)
		for j := 0; j < width; j++ {
			calculateTotalStats(ground[i][j])
			nextGenGround[i][j] = willBeAlive(CalculateAliveNeighbours(ground, i, j))
		}
	}
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

func calculateTotalStats(ground string) (totalAlive, totalDead []string) {
	if ground == "*" {
		totalAlive = append(totalAlive, "*")
	} else {
		totalDead = append(totalDead, "-")
	}
	return totalAlive, totalDead
}

func printStats(totalAlive, totalDead []string) {
	fmt.Println("Total Alive Cells = ", len(totalAlive))
	fmt.Println("Total Dead Cells = ", len(totalDead))
}

func DelayBetweenGenerations(mytime int) {
	time.Sleep(time.Duration(mytime) * time.Second)
}

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
