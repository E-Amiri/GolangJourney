package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

var (
	width  = 20
	height = 20
)

func field() {
	s := make([][]string, width, height)
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			fmt.Println(s)
		}
	}
	fmt.Println(s)
}
func initialGen() {

}

func main() {
	fmt.Println("The Width of your canvas is ", width)
	//fmt.Scanln(&width)
	fmt.Println("The Height of your canvas is ", height)
	//fmt.Scanln(&height)

	field()
}
