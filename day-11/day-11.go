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
	gameMap := setup()
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
					toChange = check1Step1(gameMap, toChange, x, y, index)
				}

				if state == occupied {
					toChange = check1Step2(gameMap, toChange, x, y, index)
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

		if len(toChange) == 0 {
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

	gameMap2 := setup()
	for z := 0; z < 500; z++ {
		toChange := make(map[string]string)

		for x := 0; x < mapX; x++ {
			for y := 0; y < mapY; y++ {
				index := fmt.Sprintf("%d:%d", x, y)
				state, ok := gameMap2[index]
				if !ok {
					log.Fatalf("%s not found", index)
				}

				if state == floor {
					continue
				}

				if state == empty {
					toChange = check2Step1(gameMap2, toChange, x, y, index)
				}

				if state == occupied {
					toChange = check2Step2(gameMap2, toChange, x, y, index)
				}
			}
		}

		for k, v := range toChange {
			_, ok := gameMap2[k]
			if !ok {
				log.Fatalf("gameMap coordinates not found: %s", k)
			}
			gameMap2[k] = v
		}

		l := len(toChange)
		if l == 0 {
			var answer int
			for _, v := range gameMap2 {
				if v == occupied {
					answer++
				}
			}
			log.Printf("day11 part 2 answer: %d", answer)
			break
		}
	}
}

func check1Step1(gameMap, toChange map[string]string, x, y int, index string) map[string]string {
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

func check1Step2(gameMap, toChange map[string]string, x, y int, index string) map[string]string {
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

func check2Step1(gameMap, toChange map[string]string, x, y int, index string) map[string]string {
	var eight int
	// Down
	xD := x + 1
	for xD := xD; xD < mapX+1; xD++ {
		if xD < mapX {
			indexD := fmt.Sprintf("%d:%d", xD, y)
			state, ok := gameMap[indexD]
			if !ok {
				log.Fatalf("down not found: %s", indexD)
			}
			if state == floor {
				continue
			}
			if state == occupied {
				break
			}
			if state == empty {
				eight++
				break
			}
		} else if xD >= mapX {
			eight++
			break
		}
	}

	// Up
	xU := x - 1
	for xU := xU; xU >= -1; xU-- {
		if xU >= 0 {
			indexU := fmt.Sprintf("%d:%d", xU, y)
			state, ok := gameMap[indexU]
			if !ok {
				log.Fatalf("up not found: %s", indexU)
			}
			if state == floor {
				continue
			}
			if state == occupied {
				break
			}
			if state == empty {
				eight++
				break
			}
		} else if xU < 0 {
			eight++
			break
		}
	}

	// Left
	yL := y - 1
	for yL := yL; yL >= -1; yL-- {
		if yL >= 0 {
			indexL := fmt.Sprintf("%d:%d", x, yL)
			state, ok := gameMap[indexL]
			if !ok {
				log.Fatalf("left not found: %s", indexL)
			}
			if state == floor {
				continue
			}
			if state == occupied {
				break
			}
			if state == empty {
				eight++
				break
			}
		} else if yL < 0 {
			eight++
			break
		}
	}

	// Right
	yR := y + 1
	for yR := yR; yR < mapY+1; yR++ {
		if yR < mapY {
			indexR := fmt.Sprintf("%d:%d", x, yR)
			state, ok := gameMap[indexR]
			if !ok {
				log.Fatalf("right not found: %s", indexR)
			}
			if state == floor {
				continue
			}
			if state == occupied {
				break
			}
			if state == empty {
				eight++
				break
			}
		} else if yR >= mapY {
			eight++
			break
		}
	}

	// Diagonal up-right
	xUR := x - 1
	yUR := y + 1
end1:
	for xUR := xUR; xUR >= -1; xUR-- {
		for yUR := yUR; yUR < mapY+1; yUR++ {
			if xUR >= 0 && yUR < mapY {
				indexUR := fmt.Sprintf("%d:%d", xUR, yUR)
				state, ok := gameMap[indexUR]
				if !ok {
					log.Fatalf("diagnoal up-righttt not found: %s", indexUR)
				}
				if state == floor {
					//continue end1
					xUR--
				}
				if state == occupied {
					break end1
				}
				if state == empty {
					eight++
					break end1
				}
			} else if xUR < 0 || yUR >= mapY {
				eight++
				break end1
			}
		}
	}

	// Diagonal up-left
	xUL := x - 1
	yUL := y - 1
end2:
	for xUL := xUL; xUL >= -1; xUL-- {
		for yUL := yUL; yUL >= -1; yUL-- {
			if xUL >= 0 && yUL >= 0 {
				indexUL := fmt.Sprintf("%d:%d", xUL, yUL)
				state, ok := gameMap[indexUL]
				if !ok {
					log.Fatalf("diagnoal up-left not found: %s", indexUL)
				}
				if state == floor {
					//continue end2
					xUL--
				}
				if state == occupied {
					break end2
				}
				if state == empty {
					eight++
					break end2
				}
			} else if xUL < 0 || yUL < 0 {
				eight++
				break end2
			}
		}
	}

	// Diagonal down-right
	xDR := x + 1
	yDR := y + 1
end3:
	for xDR := xDR; xDR < mapX+1; xDR++ {
		for yDR := yDR; yDR < mapY+1; yDR++ {
			if xDR < mapX && yDR < mapY {
				indexDR := fmt.Sprintf("%d:%d", xDR, yDR)
				state, ok := gameMap[indexDR]
				if !ok {
					log.Fatalf("diagnoal down-right not found: %s", indexDR)
				}
				if state == floor {
					xDR++
				}
				if state == occupied {
					break end3
				}
				if state == empty {
					eight++
					break end3
				}
			} else if xDR >= mapX || yDR >= mapY {
				eight++
				break end3
			}
		}
	}

	// Diagonal down-left
	xDL := x + 1
	yDL := y - 1
end4:
	for xDL := xDL; xDL < mapX+1; xDL++ {
		for yDL := yDL; yDL >= -1; yDL-- {
			if xDL < mapX && yDL >= 0 {
				indexDL := fmt.Sprintf("%d:%d", xDL, yDL)
				state, ok := gameMap[indexDL]
				if !ok {
					log.Fatalf("diagonal down-left not found: %s", indexDL)
				}
				if state == floor {
					//continue end4
					xDL++
				}
				if state == occupied {
					break end4
				}
				if state == empty {
					eight++
					break end4
				}
			} else if xDL >= mapX || yDL < 0 {
				eight++
				break end4
			}
		}
	}

	if eight == 8 {
		toChange[index] = occupied
	}

	return toChange
}

func check2Step2(gameMap, toChange map[string]string, x, y int, index string) map[string]string {
	var min int
	// Down
	xD := x + 1
	for xD := xD; xD < mapX; xD++ {
		indexD := fmt.Sprintf("%d:%d", xD, y)
		state, ok := gameMap[indexD]
		if !ok {
			log.Fatalf("down not found: %s", indexD)
		}

		if state == floor {
			continue
		}

		if state == empty {
			break
		}

		if state == occupied {
			min++
			break
		}
	}

	// Up
	xU := x - 1
	for xU := xU; xU >= 0; xU-- {
		indexU := fmt.Sprintf("%d:%d", xU, y)
		state, ok := gameMap[indexU]
		if !ok {
			log.Fatalf("up not found: %s", indexU)
		}

		if state == floor {
			continue
		}

		if state == empty {
			break
		}

		if state == occupied {
			min++
			break
		}
	}

	// Left
	yL := y - 1
	for yL := yL; yL >= 0; yL-- {
		indexL := fmt.Sprintf("%d:%d", x, yL)
		state, ok := gameMap[indexL]
		if !ok {
			log.Fatalf("left not found: %s", indexL)
		}

		if state == floor {
			continue
		}

		if state == empty {
			break
		}

		if state == occupied {
			min++
			break
		}
	}

	// Right
	yR := y + 1
	for yR := yR; yR < mapY; yR++ {
		indexR := fmt.Sprintf("%d:%d", x, yR)
		state, ok := gameMap[indexR]
		if !ok {
			log.Fatalf("right not found: %s", indexR)
		}

		if state == floor {
			continue
		}

		if state == empty {
			break
		}

		if state == occupied {
			min++
			break
		}
	}

	// Diagonal up-right
	xUR := x - 1
	yUR := y + 1
end1:
	for xUR := xUR; xUR >= 0; xUR-- {
		for yUR := yUR; yUR < mapY; yUR++ {
			if xUR >= 0 && yUR <= mapY {
				indexUR := fmt.Sprintf("%d:%d", xUR, yUR)
				state, ok := gameMap[indexUR]
				if !ok {
					log.Fatalf("up-right not found: %s", indexUR)
				}

				if state == floor {
					//continue end1
					xUR--
				}

				if state == empty {
					break end1
				}

				if state == occupied {
					min++
					break end1
				}
			}
		}
	}

	// Diagonal up-left
	xUL := x - 1
	yUL := y - 1
end2:
	for xUL := xUL; xUL >= 0; xUL-- {
		for yUL := yUL; yUL >= 0; yUL-- {
			if xUL >= 0 && yUL >= 0 {
				indexUL := fmt.Sprintf("%d:%d", xUL, yUL)
				state, ok := gameMap[indexUL]
				if !ok {
					log.Fatalf("up-left not found: %s", indexUL)
				}

				if state == floor {
					//continue end2
					xUL--
				}

				if state == empty {
					break end2
				}

				if state == occupied {
					min++
					break end2
				}
			}
		}
	}

	// Diagonal down-right
	xDR := x + 1
	yDR := y + 1
end3:
	for xDR := xDR; xDR < mapX; xDR++ {
		for yDR := yDR; yDR < mapY; yDR++ {
			if xDR < mapX && yDR <= mapY {
				indexDR := fmt.Sprintf("%d:%d", xDR, yDR)
				state, ok := gameMap[indexDR]
				if !ok {
					log.Fatalf("down-right not found: %s", indexDR)
				}

				if state == floor {
					//continue end3
					xDR++
				}

				if state == empty {
					break end3
				}

				if state == occupied {
					min++
					break end3
				}
			}
		}
	}

	// Diagonal down-left
	xDL := x + 1
	yDL := y - 1
end4:
	for xDL := xDL; xDL < mapX; xDL++ {
		for yDL := yDL; yDL >= 0; yDL-- {
			if xDL < mapX && yDL >= 0 {
				indexDL := fmt.Sprintf("%d:%d", xDL, yDL)
				state, ok := gameMap[indexDL]
				if !ok {
					log.Fatalf("down-left not found: %s", indexDL)
				}

				if state == floor {
					xDL++
					//continue end4
				}

				if state == empty {
					break end4
				}

				if state == occupied {
					min++
					break end4
				}
			}
		}
	}

	if min >= 5 {
		toChange[index] = empty
	}

	return toChange
}

func setup() map[string]string {
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

	return gameMap
}
