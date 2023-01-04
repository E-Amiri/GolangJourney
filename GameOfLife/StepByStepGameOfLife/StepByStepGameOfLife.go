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
	s := make([][]int, field.height)
	for i := 0; i < field.height; i++ {
		s[i] = make([]int, field.width)
	}
	fmt.Println(s)
}

/*
	func (field Field) setValue() bool{
		rand.Seed(time.Now().UnixNano())
		if rand.Int()

}
*/
func main() {
	field := Field{20, 10}
	Field.drawField(field)
	//field = Field{20, 30}
	//Field.drawField(field)
}
