package main

import (
	"fmt"

	"./utils"
)

func main() {
	lines := utils.GetInputLines("1.dat")
	input := utils.MapLinesToInt(lines)

	uno(input)
	dos(input)
}

func uno(input []int) {
	for i := 0; i < len(input)-1; i++ {
		for j := i + 1; j < len(input); j++ {
			if input[i]+input[j] == 2020 {
				fmt.Println("uno:", input[i]*input[j])
			}
		}
	}
}

func dos(input []int) {
	for i := 0; i < len(input)-2; i++ {
		for j := i + 1; j < len(input)-1; j++ {
			for k := j + 1; k < len(input); k++ {
				if input[i]+input[j]+input[k] == 2020 {
					fmt.Println("dos:", input[i]*input[j]*input[k])
				}
			}
		}
	}
}
