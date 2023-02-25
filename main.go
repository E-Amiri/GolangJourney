package main

import (
	"GolangJourney/GameOfLife"
	"fmt"
)

func main() {
	var height, width int
	fmt.Println("Enter the Height of Your Grid, Please ")
	fmt.Scanln(&height)
	fmt.Println("Enter the Width of Your Grid, Please ")
	fmt.Scanln(&width)

	myGrid := GameOfLife.Ground{height, width}
	ground := myGrid.initializeGround()
	GameOfLife.printGround(ground)

	for {
		ground = GameOfLife.GenerateNextGround(ground)
		GameOfLife.PrintGround(ground)
		GameOfLife.DelayBetweenGenerations(1)
		GameOfLife.ClearScreen()
	}
}
