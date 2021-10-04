package main

import (
	"bufio"
	"log"
	"os"
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
	return p.BirthYear != "" &&
		p.IssueYear != "" &&
		p.ExpirationYear != "" &&
		p.Height != "" &&
		p.HairColor != "" &&
		p.EyeColor != "" &&
		p.PassportID != ""
}
