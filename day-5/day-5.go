package main

import (
	"bufio"
	"log"
	"os"
	"sort"
)

func main() {
	f, err := os.Open("./day-5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var seats []int
	for scanner.Scan() {
		var rowMin, columnMin int
		rowMax, columnMax := 127, 7

		seat := scanner.Text()
		for _, s := range seat {
			if string(s) == "F" {
				rowMax = rowMax - ((rowMax-rowMin)+1)/2
			} else if string(s) == "B" {
				rowMin = rowMin + ((rowMax-rowMin)+1)/2
			}

			if string(s) == "L" {
				columnMax = columnMax - ((columnMax-columnMin)+1)/2
			} else if string(s) == "R" {
				columnMin = columnMin + ((columnMax-columnMin)+1)/2
			}
		}
		seats = append(seats, rowMax*8+columnMax)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Ints(seats)
	if seats == nil {
		log.Fatal("seats is nil")
	}
	log.Printf("day5 part 1 answer: %d\n", seats[len(seats)-1])
	b := 84
	for _, s := range seats {
		if s != b {
			log.Fatalf("day5 part 2 answer: %d\n", b)
		}
		b++
	}
}
