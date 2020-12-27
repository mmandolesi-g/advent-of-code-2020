package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	Byr string `json:",byr"` // Birth Year
	Iyr string `json:",iyr"` // Issue Year
	Eyr string `json:",eyr"` // Expiration Year
	Hgt string `json:",hgt"` // Height
	Hcl string `json:",hcl"` // Hair Color
	Ecl string `json:",ecl"` // Eye Color
	Pid string `json:",pid"` // Passport ID
	Cid string `json:",cid"` // Country ID
}

func main() {
	f, err := os.Open("./day-04/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var passports []passport
	var buffer string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, doWork(buffer))
			buffer = ""
			continue
		}
		buffer = fmt.Sprintf("%s %s", buffer, line)
	}
	// Catch last item in passports
	passports = append(passports, doWork(buffer))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var answer int
	for _, passport := range passports {
		if passport.Byr != "" && passport.Iyr != "" && passport.Eyr != "" && passport.Hgt != "" && passport.Hcl != "" && passport.Ecl != "" && passport.Pid != "" {
			answer++
		}
	}

	answer2 := validatePassports(passports)

	log.Printf("day4 part 1 answer: %d", answer)
	log.Printf("day4 part 2 answer: %d", answer2)
}

func validatePassports(passports []passport) int {
	var answer int
	for _, passport := range passports {
		var byrValid, iyrValid, eyrValid, hgtValid, hclValid, eclValid, pidValid bool
		if passport.Byr != "" {
			re, err := regexp.Compile("^\\d{4}$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Byr)
			if len(matches) != 1 {
				continue
			}

			byr, err := strconv.Atoi(passport.Byr)
			if err != nil {
				log.Fatal(err)
			}

			if byr >= 1920 && byr <= 2002 {
				byrValid = true
			}
		}

		if passport.Iyr != "" {
			re, err := regexp.Compile("^\\d{4}$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Iyr)
			if len(matches) != 1 {
				continue
			}

			iyr, err := strconv.Atoi(passport.Iyr)
			if err != nil {
				log.Fatal(err)
			}

			if iyr >= 2010 && iyr <= 2020 {
				iyrValid = true
			}
		}

		if passport.Eyr != "" {
			re, err := regexp.Compile("^\\d{4}$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Eyr)
			if len(matches) != 1 {
				continue
			}

			eyr, err := strconv.Atoi(passport.Eyr)
			if err != nil {
				log.Fatal(err)
			}

			if eyr >= 2020 && eyr <= 2030 {
				eyrValid = true
			}
		}

		if passport.Hgt != "" {
			re, err := regexp.Compile("(\\d+)(cm|in$)")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Hgt)
			if len(matches) != 3 {
				continue
			}

			height, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(err)
			}

			if matches[2] == "cm" {
				if height >= 150 && height <= 193 {
					hgtValid = true
				}
			} else if matches[2] == "in" {
				if height >= 59 && height <= 76 {
					hgtValid = true
				}
			}
		}

		if passport.Hcl != "" {
			re, err := regexp.Compile("^#\\w{6}$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Hcl)
			if len(matches) == 1 {
				hclValid = true
			}
		}

		if passport.Ecl != "" {
			re, err := regexp.Compile("^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Ecl)
			if len(matches) == 1 {
				eclValid = true
			}
		}

		if passport.Pid != "" {
			re, err := regexp.Compile("^\\d{9}$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Pid)
			if len(matches) == 1 {
				pidValid = true
			}
		}

		if byrValid && iyrValid && eyrValid && hgtValid && hclValid && eclValid && pidValid {
			answer++
		}
	}

	return answer
}

func doWork(passportInfo string) passport {
	var pid, hcl, byr, iyr, eyr, hgt, ecl, cid string
	parts := strings.Split(passportInfo, " ")
	for _, part := range parts {
		if strings.Contains(part, "pid") {
			pid = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "hcl") {
			hcl = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "byr") {
			byr = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "iyr") {
			iyr = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "hgt") {
			hgt = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "ecl") {
			ecl = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "eyr") {
			eyr = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "cid") {
			cid = strings.Split(part, ":")[1]
		}
	}
	return passport{
		Byr: byr,
		Iyr: iyr,
		Eyr: eyr,
		Hgt: hgt,
		Hcl: hcl,
		Ecl: ecl,
		Pid: pid,
		Cid: cid,
	}
}
