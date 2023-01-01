package main

import (
	"fmt"
	//"math/rand"
	//"time"
)

type Field struct {
	h int
	w int
}

func (field Field) makeField() {
	s := make([][]bool, field.h)
	for i := range s {
		s[i] = make([]bool, field.w)
	}
	fmt.Println(s)
}

/*func initialGen() {
	s =
}*/

func main() {
	myfield := Field{h: 20, w: 20}
	myfield.makeField()
}
