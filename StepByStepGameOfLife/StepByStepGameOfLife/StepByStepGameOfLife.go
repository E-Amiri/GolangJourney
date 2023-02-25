package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	var t []string
	var f []string
	height := 20
	width := 20
	rand.Seed(time.Now().UnixNano())

	fmt.Println("Please Enter the Height of Your Grid:")
	fmt.Scan(&height)
	fmt.Println("Please Enter the Width of Your Grid:")
	fmt.Scan(&width)

	for {
		s := make([][]string, height)
		for i := 0; i < height; i++ {
			s[i] = make([]string, width)
			fmt.Printf("\n")
			for j := 0; j < width; j++ {
				if rand.Intn(10)%2 == 0 {
					s[i][j] = "*"
					t = append(t, "*")
					fmt.Printf("%v", s[i][j])
				} else {
					s[i][j] = "-"
					f = append(f, "*")
					fmt.Printf("%v", s[i][j])
				}
			}
		}

		fmt.Println("\nTotal Alive cells =", len(t))
		fmt.Println("Total Dead cells =", len(f))
		time.Sleep(1 * time.Second)
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
