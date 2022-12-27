/*
You are making a program to analyze sport match results and
calculate the points of the given team.
The match results are stored in an array called results.
Each match has one of the following results:

"w" - won
"l" - lost
"d" - draw

A win adds three points,
a draw adds one point,
and a lost match does not add any points.

Your program needs to take the last match result
as input and append it to the results array.
After that, calculate and output the total
points the team gained from the results.
*/
package main

import "fmt"

func score(results ...string) {
	points := 0                 //total points of the Team
	for _, v := range results { // it will iterate through the results
		if v == "w" { // and add increase points based on given strings
			points += 3 // +3 for "w"(win)  and +1 for "d" (draw)
		} else if v == "d" {
			points += 1
		}
	}
	fmt.Println(points)
}
func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	var n string
	fmt.Scanln(&n)               // get the new result via input
	results = append(results, n) //add it to the "results" slice

	score(results...)
}
