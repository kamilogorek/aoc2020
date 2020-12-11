package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filename := "6.dat"
	fmt.Println("uno")
	uno(filename)
	fmt.Println("dos")
	dos(filename)
}

func uno(filename string) {
	fileHandle, _ := os.Open(filename)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	answers := make(map[string]bool)
	var total int

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if len(text) == 0 {
			total = total + len(answers)
			answers = make(map[string]bool)
			continue
		}
		for _, ch := range text {
			char := string(ch)
			answers[char] = true
		}
	}
	total = total + len(answers)

	fmt.Println(total)
}

func validate(answers map[string]int, ppl int) int {
	var total int
	for _, v := range answers {
		if v == ppl {
			total = total + 1
		}
	}
	return total
}

func dos(filename string) {
	fileHandle, _ := os.Open(filename)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)
	answers := make(map[string]int)
	var ppl int
	var total int

	for fileScanner.Scan() {
		text := fileScanner.Text()
		if len(text) == 0 {
			total = total + validate(answers, ppl)
			answers = make(map[string]int)
			ppl = 0
			continue
		}
		ppl = ppl + 1
		for _, ch := range text {
			char := string(ch)
			answers[char] = answers[char] + 1
		}
	}
	total = total + validate(answers, ppl)

	fmt.Println(total)
}
