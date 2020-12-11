package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"./utils"
)

func main() {
	lines := utils.GetInputLines("4.dat")

	var passports []string
	var passportLines []string

	for _, line := range lines {
		if len(line) == 0 {
			passports = append(passports, strings.Join(passportLines[:], " "))
			passportLines = nil
			continue
		}
		passportLines = append(passportLines, line)
	}
	passports = append(passports, strings.Join(passportLines[:], " "))

	fmt.Println("uno")
	uno(passports)
	fmt.Println("dos")
	dos(passports)
}

func isValid(passport string) bool {
	required := []string{"byr:", "iyr:", "eyr:", "hgt:", "hcl:", "ecl:", "pid:"}
	for _, req := range required {
		if !strings.Contains(passport, req) {
			return false
		}
	}
	return true
}

func isStrictValid(passport string) bool {
	for _, field := range strings.Split(passport, " ") {
		s := strings.Split(field, ":")
		key := s[0]
		val := s[1]

		if key == "byr" {
			parsedVal, _ := strconv.Atoi(val)
			if parsedVal < 1920 || parsedVal > 2002 {
				return false
			}
		}

		if key == "iyr" {
			parsedVal, _ := strconv.Atoi(val)
			if parsedVal < 2010 || parsedVal > 2020 {
				return false
			}
		}

		if key == "eyr" {
			parsedVal, _ := strconv.Atoi(val)
			if parsedVal < 2020 || parsedVal > 2030 {
				return false
			}
		}

		if key == "hgt" {
			r, _ := regexp.Compile("^(\\d+)(cm|in)$")
			match := r.FindStringSubmatch(val)
			if match == nil {
				return false
			}
			h, _ := strconv.Atoi(match[1])
			unit := match[2]

			if unit == "cm" {
				if h < 150 || h > 193 {
					return false
				}
			}

			if unit == "in" {
				if h < 59 || h > 76 {
					return false
				}
			}
		}

		if key == "hcl" {
			r, _ := regexp.Compile("^#[0-9a-f]{6}$")
			if !r.MatchString(val) {
				return false
			}
		}

		if key == "ecl" {
			r, _ := regexp.Compile("^(amb|blu|brn|gry|grn|hzl|oth)$")
			if !r.MatchString(val) {
				return false
			}

		}

		if key == "pid" {
			r, _ := regexp.Compile("^[0-9]{9}$")
			if !r.MatchString(val) {
				return false
			}
		}
	}

	return true
}

func uno(passports []string) {
	var valid int

	for _, passport := range passports {
		if isValid(passport) {
			valid = valid + 1
		}
	}

	fmt.Println(valid)
}

func dos(passports []string) {
	var valid int

	for _, passport := range passports {
		if isValid(passport) && isStrictValid(passport) {
			valid = valid + 1
		}
	}

	fmt.Println(valid)
}
