package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	passports := make([]Passport, 0)

	passport := new(Passport)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if len(scanner.Text()) == 0 {
			passports = append(passports, *passport)
			passport = new(Passport)
		}

		parts := strings.Split(scanner.Text(), " ")
		for _, part := range parts {

			attribute := strings.Split(part, ":")

			if attribute[0] == "ecl" {
				passport.EyeColor = attribute[1]
			} else if attribute[0] == "pid" {
				passport.PassportID = attribute[1]
			} else if attribute[0] == "eyr" {
				passport.ExpirationYear = attribute[1]
			} else if attribute[0] == "hcl" {
				passport.HairColor = attribute[1]
			} else if attribute[0] == "byr" {
				passport.BirthYear = attribute[1]
			} else if attribute[0] == "iyr" {
				passport.IssueYear = attribute[1]
			} else if attribute[0] == "hgt" {
				passport.Height = attribute[1]
			}
		}
	}

	cnt := 0

	for _, p := range passports {
		if p.isValid() {
			cnt++
		}
	}

	log.Print(cnt)
}

type Passport struct {
	BirthYear      string
	IssueYear      string
	ExpirationYear string
	Height         string
	HairColor      string
	EyeColor       string
	PassportID     string
}

func (p *Passport) isValid() bool {

	if p.BirthYear == "" {
		return false
	}
	byr, err := strconv.Atoi(p.BirthYear)
	if err != nil {
		panic(err)
	}
	if !(byr >= 1920 && byr <= 2002) {
		return false
	}

	if p.IssueYear == "" {
		return false
	}
	iyr, err := strconv.Atoi(p.IssueYear)
	if err != nil {
		panic(err)
	}
	if !(iyr >= 2010 && iyr <= 2020) {
		return false
	}

	if p.ExpirationYear == "" {
		return false
	}
	eyr, err := strconv.Atoi(p.ExpirationYear)
	if err != nil {
		panic(err)
	}
	if !(eyr >= 2020 && eyr <= 2030) {
		return false
	}

	if strings.HasSuffix(p.Height, "cm") {
		hgt, err := strconv.Atoi(strings.TrimSuffix(p.Height, "cm"))
		if err != nil {
			panic(err)
		}
		if !(hgt >= 150 && hgt <= 193) {
			return false
		}
	} else if strings.HasSuffix(p.Height, "in") {
		hgt, err := strconv.Atoi(strings.TrimSuffix(p.Height, "in"))
		if err != nil {
			panic(err)
		}
		if !(hgt >= 59 && hgt <= 76) {
			return false
		}
	} else {
		return false
	}

	match, _ := regexp.MatchString("^#[0-9a-f]{6}$", p.HairColor)
	if !match {
		return false
	}

	match, _ = regexp.MatchString(`^\d{9}$`, p.PassportID)
	if !match {
		return false
	}

	for _, color := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if color == p.EyeColor {
			return true
		}
	}

	return false
}
