package main

import (
	"fmt"
)

type Field struct {
	height int
	width  int
}

func (field Field) drawField() {
	s := make([][]bool, field.height)
	for i := 0; i < field.height; i++ {
		s[i] = make([]bool, field.width)
		for j := 0; j < field.width; j++ {
			if j%2 == 0 {
				s[i][j] = true
			} else {
				s[i][j] = false
			}

		}
	}
	fmt.Println(s)
}
func main() {
	field := Field{20, 10}
	Field.drawField(field)
}
