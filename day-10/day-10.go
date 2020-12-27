package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("./day-10/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var jolts []int
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		jolts = append(jolts, i)
	}
	sort.Ints(jolts)

	var diff1, diff3, totalDiff []int
	for i, j := range jolts {
		if i == 0 {
			x := jolts[i] - 0

			diff1, diff3, totalDiff = applyDiff(x, diff1, diff3, totalDiff)
		}

		// On last item, connect to my adapter which is always a difference of 3.
		if i == len(jolts)-1 {
			diff3 = append(diff3, 3)
			totalDiff = append(totalDiff, 3)
			break
		}

		x := jolts[i+1] - j
		diff1, diff3, totalDiff = applyDiff(x, diff1, diff3, totalDiff)
	}

	permutations := 1
	oneGroup := make([]int, 0)
	for _, j := range totalDiff {
		if j == 1 {
			oneGroup = append(oneGroup, j)
		} else if j == 3 {
			if len(oneGroup) >= 2 {
				if len(oneGroup) == 2 {
					permutations *= 2
				}
				if len(oneGroup) == 3 {
					permutations *= 4
				}
				if len(oneGroup) == 4 {
					permutations *= 7
				}
			}
			oneGroup = make([]int, 0)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	log.Printf("day10 part1 answer: %d\n", len(diff1)*len(diff3))
	log.Printf("day10 part2 answer: %d\n", permutations)
}

func applyDiff(x int, d1, d3, totalD []int) ([]int, []int, []int) {
	switch x {
	case 1:
		d1 = append(d1, x)
		totalD = append(totalD, x)
	case 3:
		d3 = append(d3, x)
		totalD = append(totalD, x)
	default:
		totalD = append(totalD, x) // Should never happen, differences are either 1 or 3.
	}
	return d1, d3, totalD
}
