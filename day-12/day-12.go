package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	directions := setup()

	bearing := 90  // Start facing east (090°).
	positionX := 0 // East, West
	positionY := 0 // North, South
	for _, node := range directions {
		switch node.direction {
		case "S":
			positionY = positionY - node.amount
		case "N":
			positionY = positionY + node.amount
		case "E":
			positionX = positionX + node.amount
		case "W":
			positionX = positionX - node.amount
		case "F":
			switch bearing {
			case 360, 0:
				positionY = positionY + node.amount
			case 90:
				positionX = positionX + node.amount
			case 180:
				positionY = positionY - node.amount
			case 270:
				positionX = positionX - node.amount
			}
		case "R":
			bearing = bearing + node.amount
			if bearing > 360 {
				bearing = bearing - 360
			} else if bearing < 0 {
				bearing = bearing + 360
			}
		case "L":
			bearing = bearing - node.amount
			if bearing > 360 {
				bearing = bearing - 360
			} else if bearing < 0 {
				bearing = bearing + 360
			}
		}
	}

	log.Printf("day12 part 1 answer: %d", abs(positionX)+abs(positionY))

	sx, sy := 0, 0  // Ship's x,y coordinates
	wx, wy := 10, 1 // Waypoint's x,y coordinates
	for _, node := range directions {
		switch node.direction {
		case "S":
			wy = wy - node.amount
		case "N":
			wy = wy + node.amount
		case "E":
			wx = wx + node.amount
		case "W":
			wx = wx - node.amount
		case "F":
			sx = sx + (node.amount * wx)
			sy = sy + (node.amount * wy)
		case "R":
			switch node.amount {
			case 90:
				wx, wy = rotateRight(wx, wy, 1)
			case 180:
				wx, wy = rotateRight(wx, wy, 2)
			case 270:
				wx, wy = rotateRight(wx, wy, 3)
			}
		case "L":
			switch node.amount {
			case 90:
				wx, wy = rotateLeft(wx, wy, 1)
			case 180:
				wx, wy = rotateLeft(wx, wy, 2)
			case 270:
				wx, wy = rotateLeft(wx, wy, 3)
			}
		}
	}
	log.Printf("day12 part 2 answer: %d", abs(sx)+abs(sy))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// rotateRight rotates x,y clockwise a given number of times.
func rotateRight(x, y, times int) (int, int) {
	for i := 0; i < times; i++ {
		temp := y
		y = -1 * x
		x = temp
	}
	return x, y
}

// rotateLeft rotates x,y counterclockwise a given number of times.
func rotateLeft(x, y, times int) (int, int) {
	for i := 0; i < times; i++ {
		temp := -1 * y
		y = x
		x = temp
	}
	return x, y
}

type node struct {
	direction string
	amount    int
}

func setup() []node {
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
			direction: matches[0][1],
			amount:    a,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return directions
}
