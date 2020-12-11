package main

import (
	"fmt"
	"sort"

	"./utils"
)

func main() {
	lines := utils.GetInputLines("9.dat")
	input := utils.MapLinesToInt(lines)

	uno(input)
	dos(input)
}

func uno(input []int) {
	size := 25

	for i := 0; i < len(input)-size; i++ {
		chunk := input[i : size+i]
		if !containsSum(chunk, input[size+i]) {
			fmt.Println(input[size+i])
			return
		}
	}
}

func containsSum(chunk []int, needle int) bool {
	for i := 0; i < len(chunk)-1; i++ {
		for j := i + 1; j < len(chunk); j++ {
			if chunk[i]+chunk[j] == needle {
				return true
			}
		}
	}
	return false
}

func dos(input []int) {
	// Run step 1 to get the input
	needle := 731031916

OUTER:
	for i := 0; i < len(input)-1; i++ {
		nums := []int{input[i]}
		sum := input[i]
		for j := i + 1; j < len(input); j++ {
			nums = append(nums, input[j])
			sum += input[j]
			if sum == needle {
				sort.Ints(nums)
				fmt.Println(nums)
				fmt.Println(nums[0], nums[len(nums)-1])
				fmt.Println(nums[0] + nums[len(nums)-1])
			}
			if sum > needle {
				continue OUTER
			}
		}
	}
}
