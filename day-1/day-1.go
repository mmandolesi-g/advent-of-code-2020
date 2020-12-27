package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./day-1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var num []int
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		num = append(num, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

end1:
	for _, x := range num {
		for _, y := range num {
			if x+y == 2020 {
				fmt.Printf("day1 part 1 answer: %d*%d=%d\n", x, y, x*y)
				break end1
			}
		}
	}

end2:
	for _, x := range num {
		for _, y := range num {
			for _, z := range num {
				if x+y+z == 2020 {
					fmt.Printf("day1 part 2 answer: %d*%d*%d=%d\n", x, y, z, x*y*z)
					break end2
				}
			}
		}
	}
}
