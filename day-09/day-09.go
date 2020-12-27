package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("./day-09/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var sums []int
	for scanner.Scan() {
		line := scanner.Text()
		p, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		sums = append(sums, p)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for y := 26; y < len(sums); y++ {
		var found bool
		sum := sums[y]

		preamble := sums[(y - 26):y]
		if len(preamble) != 26 {
			log.Fatalf("preamble is wrong length: %d", len(preamble))
		}
	Exit:
		for i := 0; i < len(preamble); i++ {
			for x := 0; x < len(preamble); x++ {
				if x == i {
					continue
				}

				testSum := preamble[i] + preamble[x]
				if testSum == sum {
					found = true
					break Exit
				}
			}
		}

		if !found {
			log.Printf("day9 part1 answer: %d\n", sum)
			break
		}
	}

	const invalidNumber = 57195069

End:
	for i := 0; i < len(sums); i++ {
		x := sums[i]
		var answers []int
		answers = append(answers, x)
		var totalSum int
		for z := i + 1; z < len(sums); z++ {
			s := sums[z]

			answers = append(answers, s)
			totalSum += s

			if totalSum == invalidNumber {
				sort.Ints(answers)
				log.Printf("day9 part2 answer: %d\n", answers[0]+answers[len(answers)-1])
				break End
			}

			if totalSum > invalidNumber {
				break
			}
		}
	}
}
