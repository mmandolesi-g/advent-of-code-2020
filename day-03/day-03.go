package main

import (
	"bufio"
	"log"
	"os"
)

func main() {
	answer3 := crawlRightByX(3, false)
	answer1 := crawlRightByX(1, false)
	answer5 := crawlRightByX(5, false)
	answer7 := crawlRightByX(7, false)
	answer12 := crawlRightByX(1, true)

	log.Printf("day3 part 1 answer: %d", answer3)
	log.Printf("day3 part 2 answer: %d", answer1*answer5*answer3*answer7*answer12)
}

func crawlRightByX(crawl int, twoDown bool) int {
	f, err := os.Open("./day-03/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var answer, row, index int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if row == 0 {
			row++
			continue
		}

		if twoDown {
			if row%2 != 0 {
				row++
				continue
			}
		}

		row++
		index = index + crawl

		if index > 30 {
			index = index - 31
		}

		s := scanner.Text()
		char := string(s[index])
		if char == "#" {
			answer++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return answer
}
