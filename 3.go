package main

import (
	"fmt"

	"./utils"
)

func main() {
	lines := utils.GetInputLines("3.dat")

	fmt.Println("uno")
	uno(lines)
	fmt.Println("dos")
	dos(lines)
}

func uno(lines []string) {
	trees := 0

	for i, line := range lines[1:] {
		index := ((i + 1) * 3) % len(line)
		if string(line[index]) == "#" {
			trees = trees + 1
		}
	}

	fmt.Println(trees)
}

func dos(lines []string) {
	res := 1

	trees := 0
	for i, line := range lines[1:] {
		index := (i + 1) % len(line)
		if string(line[index]) == "#" {
			trees = trees + 1
		}
	}
	res = res * trees

	trees = 0
	for i, line := range lines[1:] {
		index := ((i + 1) * 3) % len(line)
		if string(line[index]) == "#" {
			trees = trees + 1
		}
	}
	res = res * trees

	trees = 0
	for i, line := range lines[1:] {
		index := ((i + 1) * 5) % len(line)
		if string(line[index]) == "#" {
			trees = trees + 1
		}
	}
	res = res * trees

	trees = 0
	for i, line := range lines[1:] {
		index := ((i + 1) * 7) % len(line)
		if string(line[index]) == "#" {
			trees = trees + 1
		}
	}
	res = res * trees

	trees = 0
	for i, j := 2, 1; i < len(lines); i, j = i+2, j+1 {
		index := j % len(lines[i])
		if string(lines[i][index]) == "#" {
			trees = trees + 1
		}
	}
	res = res * trees

	fmt.Println(res)
}
