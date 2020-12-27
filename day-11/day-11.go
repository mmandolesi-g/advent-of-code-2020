package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	floor, empty, occupied = "floor", "empty", "occupied"
	mapX, mapY             = 97, 91
)

func main() {
	f, err := os.Open("./day-11/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	gameMap := make(map[string]string)
	var i int
	for scanner.Scan() {
		line := scanner.Text()

		for y, character := range line {
			if string(character) == "." {
				gameMap[fmt.Sprintf("%d:%d", i, y)] = floor
			} else if string(character) == "L" {
				gameMap[fmt.Sprintf("%d:%d", i, y)] = empty
			}
		}
		i++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for z := 0; z < 500; z++ {
		toChange := make(map[string]string)

		for x := 0; x < mapX; x++ {
			for y := 0; y < mapY; y++ {
				index := fmt.Sprintf("%d:%d", x, y)
				state, ok := gameMap[index]
				if !ok {
					log.Fatalf("%s not found", index)
				}

				if state == empty {
					toChange = check1(gameMap, toChange, x, y, index)
				}

				if state == occupied {
					toChange = check2(gameMap, toChange, x, y, index)
				}
			}
		}

		for k, v := range toChange {
			_, ok := gameMap[k]
			if !ok {
				log.Fatalf("gameMap coordinates not found: %s", k)
			}
			gameMap[k] = v
		}

		l := len(toChange)
		if l == 0 {
			var answer int
			for _, v := range gameMap {
				if v == occupied {
					answer++
				}
			}
			log.Printf("day11 part 1 answer: %d", answer)
			break
		}
	}
}

func check1(gameMap, toChange map[string]string, x, y int, index string) map[string]string {
	var eight int
	// Down
	xD := x + 1
	if xD < mapX {
		indexD := fmt.Sprintf("%d:%d", xD, y)
		state, ok := gameMap[indexD]
		if !ok {
			log.Fatalf("down not found: %s", indexD)
		}
		if state == empty || state == floor {
			eight++
		}
	} else if xD >= mapX {
		eight++
	}

	// Up
	xU := x - 1
	if xU >= 0 {
		indexU := fmt.Sprintf("%d:%d", xU, y)
		state, ok := gameMap[indexU]
		if !ok {
			log.Fatalf("up not found: %s", indexU)
		}
		if state == empty || state == floor {
			eight++
		}
	} else if xU < 0 {
		eight++
	}

	// Left
	yL := y - 1
	if yL >= 0 {
		indexL := fmt.Sprintf("%d:%d", x, yL)
		state, ok := gameMap[indexL]
		if !ok {
			log.Fatalf("left not found: %s", indexL)
		}
		if state == empty || state == floor {
			eight++
		}
	} else if yL < 0 {
		eight++
	}

	// Right
	yR := y + 1
	if yR < mapY {
		indexR := fmt.Sprintf("%d:%d", x, yR)
		state, ok := gameMap[indexR]
		if !ok {
			log.Fatalf("right not found: %s", indexR)
		}
		if state == empty || state == floor {
			eight++
		}
	} else if yR >= mapY {
		eight++
	}

	// Diagonal up-right
	xUR := x - 1
	yUR := y + 1
	if xUR >= 0 && yUR < mapY {
		indexUR := fmt.Sprintf("%d:%d", xUR, yUR)
		state, ok := gameMap[indexUR]
		if !ok {
			log.Fatalf("diagnoal up-right not found: %s", indexUR)
		}
		if state == empty || state == floor {
			eight++
		}
	} else if xUR < 0 || yUR >= mapY {
		eight++
	}

	// Diagonal up-left
	xUL := x - 1
	yUL := y - 1
	if xUL >= 0 && yUL >= 0 {
		indexUL := fmt.Sprintf("%d:%d", xUL, yUL)
		state, ok := gameMap[indexUL]
		if !ok {
			log.Fatalf("diagnoal up-left not found: %s", indexUL)
		}
		if state == empty || state == floor {
			eight++
		}
	} else if xUL < 0 || yUL < 0 {
		eight++
	}

	// Diagonal down-right
	xDR := x + 1
	yDR := y + 1
	if xDR < mapX && yDR < mapY {
		indexDR := fmt.Sprintf("%d:%d", xDR, yDR)
		state, ok := gameMap[indexDR]
		if !ok {
			log.Fatalf("diagnoal down-right not found: %s", indexDR)
		}
		if state == empty || state == floor {
			eight++
		}
	} else if xDR >= mapX || yDR >= mapY {
		eight++
	}

	// Diagonal down-left
	xDL := x + 1
	yDL := y - 1
	if xDL < mapX && yDL >= 0 {
		indexDL := fmt.Sprintf("%d:%d", xDL, yDL)
		state, ok := gameMap[indexDL]
		if !ok {
			log.Fatalf("diagonal down-left not found: %s", indexDL)
		}
		if state == empty || state == floor {
			eight++
		}
	} else if xDL >= mapX || yDL < 0 {
		eight++
	}

	if eight == 8 {
		toChange[index] = occupied
	}

	return toChange
}

func check2(gameMap, toChange map[string]string, x, y int, index string) map[string]string {
	var minFour int
	// Down
	xD := x + 1
	if xD <= mapX {
		indexD := fmt.Sprintf("%d:%d", xD, y)
		state := gameMap[indexD]
		if state == occupied {
			minFour++
		}
	}

	// Up
	xU := x - 1
	if xU >= 0 {
		indexU := fmt.Sprintf("%d:%d", xU, y)
		state := gameMap[indexU]
		if state == occupied {
			minFour++
		}
	}

	// Left
	yL := y - 1
	if yL >= 0 {
		indexL := fmt.Sprintf("%d:%d", x, yL)
		state := gameMap[indexL]
		if state == occupied {
			minFour++
		}
	}

	// Right
	yR := y + 1
	if yR <= mapY {
		indexR := fmt.Sprintf("%d:%d", x, yR)
		state := gameMap[indexR]
		if state == occupied {
			minFour++
		}
	}

	// Diagonal up-right
	xUR := x - 1
	yUR := y + 1
	if xUR >= 0 && yUR <= mapY {
		indexUR := fmt.Sprintf("%d:%d", xUR, yUR)
		state := gameMap[indexUR]
		if state == occupied {
			minFour++
		}
	}

	// Diagonal up-left
	xUL := x - 1
	yUL := y - 1
	if xUL >= 0 && yUL >= 0 {
		indexUL := fmt.Sprintf("%d:%d", xUL, yUL)
		state := gameMap[indexUL]
		if state == occupied {
			minFour++
		}
	}

	// Diagonal down-right
	xDR := x + 1
	yDR := y + 1
	if xDR <= mapX && yDR <= mapY {
		indexDR := fmt.Sprintf("%d:%d", xDR, yDR)
		state := gameMap[indexDR]
		if state == occupied {
			minFour++
		}
	}

	// Diagonal down-left
	xDL := x + 1
	yDL := y - 1
	if xDL <= mapX && yDL >= 0 {
		indexDL := fmt.Sprintf("%d:%d", xDL, yDL)
		state := gameMap[indexDL]
		if state == occupied {
			minFour++
		}
	}

	if minFour >= 4 {
		toChange[index] = empty
	}

	return toChange
}
