package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	type password struct {
		Min    int    `json:",min"`
		Max    int    `json:",max"`
		Letter string `json:",letter"`
		Value  string `json:",value"`
	}

	f, err := os.Open("./day-02/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var passwords []password
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, " ")

		minMax := strings.Split(parts[0], "-")
		letter := string(parts[1][0])

		min, err := strconv.Atoi(minMax[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			log.Fatal(err)
		}

		passwords = append(passwords, password{
			Min:    min,
			Max:    max,
			Letter: letter,
			Value:  parts[2],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var answer int
	for _, p := range passwords {
		var numFound int
		for _, s := range p.Value {
			if p.Letter == string(s) {
				numFound++
			}
		}

		if numFound >= p.Min && numFound <= p.Max {
			answer++
		}
	}

	log.Printf("day2 part 1 answer: %d\n", answer)

	var answer2 int
	for _, p := range passwords {
		char1 := string(p.Value[p.Min-1])
		char2 := string(p.Value[p.Max-1])

		if (p.Letter == char1 && p.Letter != char2) || (p.Letter != char1 && p.Letter == char2) {
			answer2++
			continue
		}
	}

	log.Printf("day2 part 2 answer: %d\n", answer2)
}
