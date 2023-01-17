package main

import (
	"fmt"
	"math/rand"

	//"os"
	//"os/exec"
	"time"
)

func main() {
	height := 0
	width := 0

	fmt.Println("Enter the Height of Your Grid, Please ")
	fmt.Scanln(&height)
	fmt.Println("Enter the Width of Your Grid, Please ")
	fmt.Scanln(&width)

	s := make([][]string, height)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < height; i++ {
		s[i] = make([]string, width)
		fmt.Println("\n")
		for j := 0; j < width; j++ {
			if rand.Intn(10)%2 == 0 {
				s[i][j] = "*"
				fmt.Printf("%v ", s[i][j])
			} else {
				s[i][j] = "-"
				fmt.Printf("%v ", s[i][j])
			}
		}
	}
	for {
		fmt.Println("\n_______________________")
		alive := 0
		thisAlive := 0
		for i := 1; i < height; i++ {
			fmt.Println("\n")
			for j := 1; j < width; j++ {
				for di := i - 1; di <= i; di++ {
					for dj := j - 1; dj <= j; dj++ {
						if s[di][dj] == "*" && di != i && dj != j {
							alive++
						} else if di != i && dj != j {
							thisAlive = 1
						}
					}
				}
				if alive < 4 && alive > 1 && thisAlive == 1 || alive == 3 && thisAlive != 1 {
					s[i][j] = "*"
					fmt.Printf("%v ", s[i][j])
				} else {
					s[i][j] = "-"
					fmt.Printf("%v ", s[i][j])
				}
			}
		}
		/*	time.Sleep(1 * time.Second)
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()*/
	}

}
