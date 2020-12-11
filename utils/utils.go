package utils

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

// GetInputLines | filename string -> lines []string
func GetInputLines(filename string) []string {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(raw), "\n")
	if len(lines[len(lines)-1]) == 0 {
		return lines[:len(lines)-1]
	}
	return lines
}

// MapLinesToInt | lines []string -> lines []int
func MapLinesToInt(lines []string) []int {
	mapped := make([]int, len(lines))
	for i, v := range lines {
		mapped[i], _ = strconv.Atoi(v)
	}
	return mapped
}

// GetRegexpGroups | regex, url string -> grou[s map[string]string
func GetRegexpGroups(regex, url string) map[string]string {
	r := regexp.MustCompile(regex)
	match := r.FindStringSubmatch(url)
	groups := make(map[string]string)
	for i, name := range r.SubexpNames() {
		if i > 0 && i <= len(match) {
			groups[name] = match[i]
		}
	}
	return groups
}

// ArraysEqual | a, b []interface{} -> bool
func ArraysEqual(a, b []string) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
