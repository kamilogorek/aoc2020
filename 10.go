package main

import (
	"fmt"
	"sort"

	"./utils"
)

func main() {
	lines := utils.GetInputLines("10.dat")
	input := utils.MapLinesToInt(lines)

	uno(input)
	dos(input)
}

func uno(input []int) {
	plugs := input
	sort.Ints(plugs)

	ones := 1
	threes := 1

	for i := 0; i < len(plugs)-1; i++ {
		diff := plugs[i+1] - plugs[i]
		if diff == 1 {
			ones++
		}
		if diff == 3 {
			threes++
		}
	}

	fmt.Println(ones * threes)
}

func dos(input []int) {
	plugs := input
	sort.Ints(plugs)
	plugs = append([]int{0}, plugs...)
	plugs = append(plugs, plugs[len(plugs)-1]+3)

	// Answer: 19208
	var groups []int
	var diffs []int
	acc := 0

	for i := 0; i < len(plugs)-1; i++ {
		diff := plugs[i+1] - plugs[i]
		diffs = append(diffs, diff)
		if diff == 1 {
			acc++
		} else {
			if acc > 1 {
				groups = append(groups, acc)
			}
			acc = 0
		}
	}

	// 4 = 7
	// 3 = 4
	// 2 = 2
	// 1 = 1

	res := 1

	for _, v := range groups {
		if v == 4 {
			res *= 7
		}
		if v == 3 {
			res *= 4
		}
		if v == 2 {
			res *= 2
		}
	}
	// 7/7/4/2/7/7
	// 4/4/3/2/4/4
	fmt.Println(groups, acc, res)
}
