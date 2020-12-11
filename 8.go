package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	filename := "8.dat"
	fmt.Println("uno")
	uno(filename)
	fmt.Println("dos")
	dos(filename)
}

func uno(filename string) {
	fileHandle, _ := os.Open(filename)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	var instructions [][]string

	for fileScanner.Scan() {
		text := fileScanner.Text()
		r, _ := regexp.Compile("([a-z]{3}) ([-+]\\d+)")
		match := r.FindStringSubmatch(text)
		instructions = append(instructions, match[1:])
	}

	var acc int
	var visited []int

	for i := 0; i < len(instructions); {
		for _, v := range visited {
			if v == i {
				fmt.Println(acc)
				return
			}
		}
		visited = append(visited, i)
		instruction := instructions[i][0]
		value, _ := strconv.Atoi(instructions[i][1])
		if instruction == "nop" {
			i++
			continue
		}
		if instruction == "acc" {
			acc += value
			i++
			continue
		}
		if instruction == "jmp" {
			i += value
			continue
		}
	}

	fmt.Println(instructions)
}

func dos(filename string) {
	fileHandle, _ := os.Open(filename)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	var instructions [][]string

	for fileScanner.Scan() {
		text := fileScanner.Text()
		r, _ := regexp.Compile("([a-z]{3}) ([-+]\\d+)")
		match := r.FindStringSubmatch(text)
		instructions = append(instructions, match[1:])
	}

MODIFIER:
	for i := 0; i < len(instructions); i++ {
		modified := make([][]string, len(instructions))

		for i := range instructions {
			modifiedInstruction := make([]string, 2)
			copy(modifiedInstruction, instructions[i])
			modified[i] = modifiedInstruction
		}

		if modified[i][0] == "acc" {
			continue
		}

		if modified[i][0] == "nop" {
			modified[i][0] = "jmp"
		} else if modified[i][0] == "jmp" {
			modified[i][0] = "nop"
		}

		var acc int
		var visited []int

		for j := 0; j < len(modified); {
			for _, v := range visited {
				if v == j {
					continue MODIFIER
				}
			}
			visited = append(visited, j)
			instruction := modified[j][0]
			value, _ := strconv.Atoi(modified[j][1])
			if instruction == "nop" {
				j++
				continue
			}
			if instruction == "acc" {
				acc += value
				j++
				continue
			}
			if instruction == "jmp" {
				if value == 0 {
					continue MODIFIER
				}
				j += value
				continue
			}
		}

		fmt.Println(acc)
	}
}
