package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Step struct {
	Action string
	Amount int
}

func main() {
	f, err := os.Open("./day-08/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var steps []Step
	var num int
	for scanner.Scan() {
		num++
		line := scanner.Text()

		lineParsed := strings.Split(line, " ")
		action := lineParsed[0]
		amount := lineParsed[1]

		a, err := strconv.Atoi(amount)
		if err != nil {
			log.Fatal(err)
		}

		s := Step{
			Action: action,
			Amount: a,
		}

		steps = append(steps, s)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	changeMap := make(map[int]Step)
	for i, step := range steps {

		for k, v := range changeMap {
			steps[k] = v
		}

		if step.Action == "jmp" {
			steps[i].Action = "nop"
			changeMap[i] = step
		}

		_ = stepThrough(steps)
	}

	changeMap2 := make(map[int]Step)
	for i, step := range steps {

		for k, v := range changeMap2 {
			steps[k] = v
		}

		if step.Action == "nop" {
			steps[i].Action = "jmp"
			changeMap2[i] = step
		}

		_ = stepThrough(steps)
	}

	log.Printf("day8 part1 answer: %d\n", 1)
}

func stepThrough(steps []Step) int {
	var accumulator int
	m := make(map[int]Step)
	var mappedSteps []Step

	for i := 0; i < len(steps); i++ {
		if i >= 600 {
			log.Printf("day8 part2 answer: %d", accumulator)
		}

		_, ok := m[i]
		if ok {
			mappedSteps = append(mappedSteps, steps[i])
			//log.Println("break")
			break
		}

		if steps[i].Action == "acc" {
			accumulator = accumulator + steps[i].Amount
			m[i] = steps[i]
			mappedSteps = append(mappedSteps, steps[i])
			continue
		}
		if steps[i].Action == "jmp" {
			i = (i + steps[i].Amount) - 1
			m[i] = steps[i]
			mappedSteps = append(mappedSteps, steps[i])
			continue
		}
		if steps[i].Action == "nop" {
			m[i] = steps[i]
			mappedSteps = append(mappedSteps, steps[i])
			continue
		}
	}

	return accumulator
}
