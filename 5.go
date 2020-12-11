package main

import (
	"fmt"
	"sort"

	"./utils"
)

func main() {
	lines := utils.GetInputLines("5.dat")

	fmt.Println("uno")
	uno(lines)
	fmt.Println("dos")
	dos(lines)
}

func getSeatId(line string) int {
	top := 0
	bottom := 128
	left := 0
	right := 8
	for _, ch := range line {
		char := string(ch)

		if char == "F" {
			diff := bottom - top
			bottom = bottom - diff/2
		}

		if char == "B" {
			diff := bottom - top
			top = top + diff/2
		}

		if char == "L" {
			diff := right - left
			right = right - diff/2
		}

		if char == "R" {
			diff := right - left
			left = left + diff/2
		}
	}
	return top*8 + left
}

func uno(lines []string) {
	high := 0

	for _, line := range lines {
		id := getSeatId(line)
		if id > high {
			high = id
		}
	}

	fmt.Println(high)
}

func dos(lines []string) {
	var seats []int

	for _, line := range lines {
		id := getSeatId(line)
		seats = append(seats, id)
	}
	sort.Ints(seats)
	fmt.Println(seats)

	var prev int
	for _, curr := range seats {
		if prev == 0 {
			prev = curr
			continue
		}
		if prev == curr-1 {
			prev = curr
			continue
		}
		fmt.Println(prev + 1)
		return
	}
}
