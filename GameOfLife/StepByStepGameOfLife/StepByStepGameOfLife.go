package main

import (
	"fmt"
	//"math/rand"
)

type Field struct {
	height int
	width  int
}

func (field Field) drawField() {
	s := make([][]int, field.height)
	for i := 0; i < field.height; i++ {
		s[i] = make([]int, field.width)
	}
	fmt.Println(s)
}

func main() {
	field := Field{20, 10}
	Field.drawField(field)
}
