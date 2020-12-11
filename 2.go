package main

import (
	"fmt"
	"strconv"
	"strings"

	"./utils"
)

func main() {
	lines := utils.GetInputLines("2.dat")
	uno(lines)
	dos(lines)
}

func uno(lines []string) {
	valid := 0
	for _, line := range lines {
		x := utils.GetRegexpGroups("(?P<from>\\d+)-(?P<to>\\d+) (?P<letter>[a-z]+): (?P<entry>[a-z]+)", line)
		c := strings.Count(x["entry"], x["letter"])
		from, _ := strconv.Atoi(x["from"])
		to, _ := strconv.Atoi(x["to"])
		if c >= from && c <= to {
			valid = valid + 1
		}
	}
	fmt.Println(valid)
}

func dos(lines []string) {
	valid := 0
	for _, line := range lines {
		found := 0
		x := utils.GetRegexpGroups("(?P<from>\\d+)-(?P<to>\\d+) (?P<letter>[a-z]+): (?P<entry>[a-z]+)", line)
		pos1, _ := strconv.Atoi(x["from"])
		pos2, _ := strconv.Atoi(x["to"])
		if string(x["entry"][pos1-1]) == x["letter"] {
			found = found + 1
		}
		if string(x["entry"][pos2-1]) == x["letter"] {
			found = found + 1
		}
		if found == 1 {
			valid = valid + 1
		}
	}
	fmt.Println(valid)
}
