package main

import (
	"fmt"
	"strings"

	"./utils"
)

func main() {
	lines := utils.GetInputLines("11-example.dat")

	uno(lines)
	dos(lines)
}

func uno(lines []string) {
	var input [][]string

	for _, v := range lines {
		input = append(input, strings.Split(v, ""))
	}

	//evolved := evolve([][]string{})
	evolved := evolve(input)

//	var x [][]string
//	var y [][]string
//
//	x = append(x, []string{".", ".", "."})
//	x = append(x, []string{".", ".", "."})
//	x = append(x, []string{".", "x", "."})
//	y = append(y, []string{".", ".", "."})
//	y = append(y, []string{".", ".", "."})
//	y = append(y, []string{".", "x", "."})
//
//	fmt.Println(same(x, y))

	var occupied int
	for i := 0; i < len(evolved); i++ {
		for j := 0; j < len(evolved[i]); j++ {
			if evolved[i][j] == "#" {
				occupied++
			}
		}
	}

	fmt.Println(occupied)
}

func same(x [][]string, y [][]string) bool {
	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			if x[i][j] != y[i][j] {
				return false
			}
		}
	}
	return true

}

func evolve(input [][]string) [][]string {
	seats := input

	var post [][]string
	for i := 0; i < len(seats); i++ {
		post = append(post, make([]string, len(seats[i])))
		for j := 0; j < len(seats[i]); j++ {
			var batch [][]string

			// first line
			if i == 0 {
				batch = append(batch, []string{".", ".", "."})
			} else {
				if j == 0 {
					batch = append(batch, []string{".", seats[i][j], seats[i][j+1]})
				} else if j == len(seats[i]) - 1 {
					batch = append(batch, []string{seats[i][j-1], seats[i][j], "."})
				} else {
					batch = append(batch, []string{seats[i][j-1], seats[i][j], seats[i][j+1]})
				}
			}

			if j == 0 {
				batch = append(batch, []string{".", seats[i][j], seats[i][j+1]})
			} else if j == len(seats[i]) - 1 {
				batch = append(batch, []string{seats[i][j-1], seats[i][j], "."})
			} else {
				batch = append(batch, []string{seats[i][j-1], seats[i][j], seats[i][j+1]})
			}

			// last line
			if i == len(seats) - 1 {
				batch = append(batch, []string{".", ".", "."})
			} else {
				if j == 0 {
					batch = append(batch, []string{".", seats[i+1][j], seats[i+1][j+1]})
				} else if j == len(seats[i]) - 1 {
					batch = append(batch, []string{seats[i+1][j-1], seats[i+1][j], "."})
				} else {
					batch = append(batch, []string{seats[i+1][j-1], seats[i+1][j], seats[i+1][j+1]})
				}
			}

			n := collectNeighbors(batch)
			fmt.Println(i, j, batch, n)

			if n ==0 {
				post[i][j] = "#"
			}
			if n >= 4 {
				post[i][j] = "L"
			}
		}
	}

	fmt.Println("====")
	fmt.Println(seats)
	fmt.Println(post)

	if same(seats, post) {
		return post
	} else {
		return evolve(post)
	}
}

func collectNeighbors(seats [][]string) int {
	var neighbors int
	if string(seats[0][0]) == "#" {
		neighbors++
	}
	if string(seats[0][1]) == "#" {
		neighbors++
	}
	if string(seats[0][2]) == "#" {
		neighbors++
	}
	if string(seats[1][0]) == "#" {
		neighbors++
	}
	if string(seats[1][2]) == "#" {
		neighbors++
	}
	if string(seats[2][0]) == "#" {
		neighbors++
	}
	if string(seats[2][1]) == "#" {
		neighbors++
	}
	if string(seats[2][2]) == "#" {
		neighbors++
	}
	return neighbors
}

func dos(input []string) {
	fmt.Println(len(input))
}
