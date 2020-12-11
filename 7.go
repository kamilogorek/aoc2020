package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/kr/pretty"
)

func main() {
	filename := "7.dat"
	fmt.Println("uno")
	uno(filename)
	fmt.Println("dos")
	dos(filename)
}

func uno(filename string) {
	fileHandle, _ := os.Open(filename)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	nodes := make(map[string][]string)

	for fileScanner.Scan() {
		text := fileScanner.Text()
		r, _ := regexp.Compile("([a-z]{3,} [a-z]+) bag")
		match := r.FindAllStringSubmatch(text, -1)

		var root string
		for i, m := range match {
			color := string(m[1])
			if i == 0 {
				root = color
				continue
			}
			nodes[root] = append(nodes[root], color)
		}
	}

	var matches []string
	for k := range nodes {
		if canAccess(nodes, k, "shiny gold", []string{k}) {
			matches = append(matches, k)
		}
	}

	fmt.Println(len(matches))
}

func canAccess(nodes map[string][]string, key string, needle string, visited []string) bool {
	node := nodes[key]

	for _, k := range node {
		if k == needle {
			return true
		}

		for _, v := range visited {
			if v == k {
				return false
			}
		}

		if canAccess(nodes, k, needle, append(visited, k)) {
			return true
		}
	}

	return false
}

func dos(filename string) {
	fileHandle, _ := os.Open(filename)
	defer fileHandle.Close()
	fileScanner := bufio.NewScanner(fileHandle)

	rawColors := make(map[string]string)
	for fileScanner.Scan() {
		text := fileScanner.Text()
		r, _ := regexp.Compile("([a-z]{3,} [a-z]+) bag")
		match := r.FindStringSubmatch(text)
		rawColors[match[1]] = text
	}

	colors := make(map[string]map[string]int)
	for k, color := range rawColors {
		r, _ := regexp.Compile("(\\d+) ([a-z]{3,} [a-z]+) bag")
		match := r.FindAllStringSubmatch(color, -1)
		if colors[k] == nil {
			colors[k] = make(map[string]int)
		}
		if match == nil {
			continue
		}
		for _, v := range match {
			count, _ := strconv.Atoi(v[1])
			col := v[2]
			colors[k][col] = count
		}
	}

	// Initial iteration shouldnt add 1 because the root bag itself doesnt count
	pretty.Log(sumChildren(colors, "shiny gold") - 1)
}

func sumChildren(colors map[string]map[string]int, key string) int {
	children := colors[key]
	if len(children) == 0 {
		return 1
	}

	total := 0
	for k, v := range children {
		fmt.Println(k, sumChildren(colors, k))
		total = total + sumChildren(colors, k)*v
	}
	// Add the bag that contains children itself
	return total + 1
}
