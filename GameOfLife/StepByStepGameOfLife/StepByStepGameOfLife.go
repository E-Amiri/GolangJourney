package main

import (
	"fmt"
)

type Field struct { // defines a Field type
	height int //in which we declare
	width  int // size of the grid
}

func (field Field) drawField() { // drawField method of the Field structs
	s := make([][]bool, field.height)   // generates the grid through a 2dinmetional slice
	for i := 0; i < field.height; i++ { // we nest a slice inside another slice via iterating
		s[i] = make([]bool, field.width)   // i index refers to the height of the grid
		for j := 0; j < field.width; j++ { // j index refers to the width of the grid
			if j%2 == 0 { // set the grid values based on modulus of the index
				s[i][j] = true
			} else {
				s[i][j] = false
			}

		}
	}
	fmt.Println(s)
}
func main() {
	field := Field{20, 10} // set the size of the grid via parameters
	Field.drawField(field)
}
