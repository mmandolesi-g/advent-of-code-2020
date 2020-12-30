package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	f, err := os.Open("./day-12/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var directions []node
	re := regexp.MustCompile(`(\w)(\d+)`)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, 1)

		a, err := strconv.Atoi(matches[0][2])
		if err != nil {
			log.Fatal(err)
		}
		directions = append(directions, node{
			Direction: matches[0][1],
			Amount:    a,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	bearing := 90  // Start facing east (090°).
	positionX := 0 // East, West
	positionY := 0 // North, South
	for _, node := range directions {
		switch node.Direction {
		case "S":
			positionY = positionY - node.Amount
		case "N":
			positionY = positionY + node.Amount
		case "E":
			positionX = positionX + node.Amount
		case "W":
			positionX = positionX - node.Amount
		case "F":
			switch bearing {
			case 360, 0:
				positionY = positionY + node.Amount
			case 90:
				positionX = positionX + node.Amount
			case 180:
				positionY = positionY - node.Amount
			case 270:
				positionX = positionX - node.Amount
			}
		case "R":
			bearing = bearing + node.Amount
			if bearing > 360 {
				bearing = bearing - 360
			} else if bearing < 0 {
				bearing = bearing + 360
			}
		case "L":
			bearing = bearing - node.Amount
			if bearing > 360 {
				bearing = bearing - 360
			} else if bearing < 0 {
				bearing = bearing + 360
			}
		}
	}

	log.Printf("day11 part 1 answer: %d", abs(positionX)+abs(positionY))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type node struct {
	Direction string
	Amount    int
}
