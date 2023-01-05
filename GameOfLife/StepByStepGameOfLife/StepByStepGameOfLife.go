package main

import (
	"fmt"
)

type Field struct {
	height int
	width  int
	//mainSlice [][]bool
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

/*
	func (field Field) setValue(x, y int, b bool) [][]bool {
		field.mainSlice[x][y] = b
		return field.mainSlice
	}
*/
func main() {
	//var f Field
	//var s [][]bool = Field.setValue(f, 20, 10, true)
	field := Field{20, 10}
	Field.drawField(field)
	//field = Field{20, 30}
	//Field.drawField(field)
}
