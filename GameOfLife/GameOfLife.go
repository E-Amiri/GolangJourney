package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

var (
	width  = 20
	height = 20
	//s      = make([][]string, width, height)
)

func field() {
	s := make([][]bool, width)
	for i := range s {
		s[i] = make([]bool, height)
	}
	fmt.Println(s)
}

/*func initialGen() {
	s =
}*/

func main() {
	fmt.Println("The Width of your canvas is ", width)
	//fmt.Scanln(&width)
	fmt.Println("The Height of your canvas is ", height)
	//fmt.Scanln(&height)

	field()
}
