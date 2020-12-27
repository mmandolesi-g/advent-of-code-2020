package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//day1()
	//day2()
	//day3()
	//day4()
	//day5()
	//day6()
	//day6Part2()
	//day7()
	//day8()
	//day9()
	//day10()
	day11()
}

func day11() {
	const (
		floor, empty, occupied = "floor", "empty", "occupied"
		mapX, mapY             = 97, 91
	)

	f, err := os.Open("./inputs/day11.txt")
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
				}

				if state == occupied {
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
				if v == "occupied" {
					answer++
				}
			}
			log.Printf("day11 part1 answer: %d\n", answer)
			break
		}
	}
}

func day10() {
	f, err := os.Open("./inputs/day10.txt")
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

func day9() {
	f, err := os.Open("./inputs/day9.txt")
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

func day8() {
	f, err := os.Open("./inputs/day8.txt")
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

type Step struct {
	Action string
	Amount int
}

func day7() {
	f, err := os.Open("./inputs/day7.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	re := regexp.MustCompile(`((\d*\s*)(\w+\s\w+\s)bag(s)*)`)
	//re := regexp.MustCompile(`((\w+\s\w+\s)bag(s)*)`)
	bags := make(map[string]*Node)
	for scanner.Scan() {
		line := scanner.Text()

		matches := re.FindAllStringSubmatch(line, 10)
		node := &Node{}
		for i, m := range matches {
			if i == 0 {
				bag, ok := bags[m[3]]
				if !ok {
					bags[m[3]] = node
					node.Name = m[3]
				} else {
					node = bag
				}
				continue
			}
			bag, ok := bags[m[3]]
			var nx *Node
			if !ok {
				nx = &Node{}
			} else {
				nx = bag
			}

			if m[3] == "no other " {
				continue
			}

			nx.Name = m[3]
			x, err := strconv.Atoi(strings.TrimSpace(m[2]))
			if err != nil {
				log.Fatal(err)
			}
			nx.Parents = append(nx.Parents, node)
			if node.Children == nil {
				m := make(map[*Node]int)
				node.Children = m
			}
			node.Children[nx] = x
			bags[m[3]] = nx
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	b, ok := bags["shiny gold "]
	if !ok {
		log.Fatal("bag not found")
	}

	outer := make(map[string]*Node)
	log.Printf("day7 part1 answer: %d\n", len(recursiveCount(b, outer)))

	var answer2 int
	answer2 = recursiveCount2(b, 1, 0)
	log.Printf("day7 part2 answer: %d\n", answer2)
}

type Node struct {
	Name     string
	Children map[*Node]int
	Parents  []*Node
}

func recursiveCount(node *Node, outer map[string]*Node) map[string]*Node {
	for _, n := range node.Parents {
		outer[n.Name] = n
		if n.Parents != nil {
			outer = recursiveCount(n, outer)
		}
	}
	return outer
}

func recursiveCount2(parent *Node, multiplier int, total int) int {
	for child, count := range parent.Children {
		newmul := multiplier * count
		total += newmul
		if child.Children != nil {
			total = recursiveCount2(child, newmul, total)
		}
	}
	return total
}

func day6Part2() {
	f, err := os.Open("./inputs/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var answers []int
	var buffer string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			answers = append(answers, process62(strings.TrimSpace(buffer)))
			buffer = ""
			continue
		}
		buffer = fmt.Sprintf("%s %s", buffer, line)
	}
	// Catch last item
	answers = append(answers, process62(strings.TrimSpace(buffer)))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var answer int
	for _, i := range answers {
		answer = i + answer
	}

	log.Printf("day6 part 2 answer: %d\n", answer)

}

func process62(data string) int {
	var alphabet = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	groupResponse := strings.Split(data, " ")
	var answer int

	for _, letter := range alphabet {
		var groupAnswers []string

		for _, response := range groupResponse {
			if strings.Contains(response, letter) {
				groupAnswers = append(groupAnswers, letter)
			}
		}

		if len(groupAnswers) == len(groupResponse) {
			answer++
		}
	}

	return answer
}

func day6() {
	f, err := os.Open("./inputs/day6.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var answers []int
	var buffer string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			answers = append(answers, process61(buffer))
			buffer = ""
			continue
		}
		buffer = fmt.Sprintf("%s%s", buffer, line)
	}
	// Catch last item
	answers = append(answers, process61(buffer))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var answer1 int
	for _, i := range answers {
		answer1 = i + answer1
	}

	log.Printf("day6 part 1 answer: %d\n", answer1)
}

func process61(data string) int {
	m := make(map[string]bool)
	for _, item := range data {
		if _, ok := m[string(item)]; !ok {
			m[string(item)] = true
		}
	}

	var result []string
	for item, _ := range m {
		result = append(result, item)
	}
	return len(strings.Join(result, ""))
}

func day5() {
	f, err := os.Open("day5.txt")
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

	sort.Ints(seats)
	log.Printf("day5 part 1 answer: %d\n", seats[len(seats)-1])
	b := 84
	for _, s := range seats {
		if s != b {
			log.Fatalf("day5 part 2 answer: %d\n", b)
		}
		b++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func day4() {
	f, err := os.Open("day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var passports []passport
	var buffer string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, doWork(buffer))
			buffer = ""
			continue
		}
		buffer = fmt.Sprintf("%s %s", buffer, line)
	}
	// Catch last item in passports
	passports = append(passports, doWork(buffer))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var answer int
	for _, passport := range passports {
		if passport.Byr != "" && passport.Iyr != "" && passport.Eyr != "" && passport.Hgt != "" && passport.Hcl != "" && passport.Ecl != "" && passport.Pid != "" {
			answer++
		}
	}

	answer2 := validatePassports(passports)

	fmt.Printf("day4 part 1 answer: %d\n", answer)
	fmt.Printf("day4 part 2 answer: %d\n", answer2)
}

type passport struct {
	Byr string `json:",byr"` // Birth Year
	Iyr string `json:",iyr"` // Issue Year
	Eyr string `json:",eyr"` // Expiration Year
	Hgt string `json:",hgt"` // Height
	Hcl string `json:",hcl"` // Hair Color
	Ecl string `json:",ecl"` // Eye Color
	Pid string `json:",pid"` // Passport ID
	Cid string `json:",cid"` // Country ID
}

func validatePassports(passports []passport) int {
	var answer int
	for _, passport := range passports {
		var byrValid, iyrValid, eyrValid, hgtValid, hclValid, eclValid, pidValid bool
		if passport.Byr != "" {
			re, err := regexp.Compile("^\\d{4}$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Byr)
			if len(matches) != 1 {
				continue
			}

			byr, err := strconv.Atoi(passport.Byr)
			if err != nil {
				log.Fatal(err)
			}

			if byr >= 1920 && byr <= 2002 {
				byrValid = true
			}
		}

		if passport.Iyr != "" {
			re, err := regexp.Compile("^\\d{4}$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Iyr)
			if len(matches) != 1 {
				continue
			}

			iyr, err := strconv.Atoi(passport.Iyr)
			if err != nil {
				log.Fatal(err)
			}

			if iyr >= 2010 && iyr <= 2020 {
				iyrValid = true
			}
		}

		if passport.Eyr != "" {
			re, err := regexp.Compile("^\\d{4}$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Eyr)
			if len(matches) != 1 {
				continue
			}

			eyr, err := strconv.Atoi(passport.Eyr)
			if err != nil {
				log.Fatal(err)
			}

			if eyr >= 2020 && eyr <= 2030 {
				eyrValid = true
			}
		}

		if passport.Hgt != "" {
			re, err := regexp.Compile("(\\d+)(cm|in$)")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Hgt)
			if len(matches) != 3 {
				continue
			}

			height, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Fatal(err)
			}

			if matches[2] == "cm" {
				if height >= 150 && height <= 193 {
					hgtValid = true
				}
			} else if matches[2] == "in" {
				if height >= 59 && height <= 76 {
					hgtValid = true
				}
			}
		}

		if passport.Hcl != "" {
			re, err := regexp.Compile("^#\\w{6}$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Hcl)
			if len(matches) == 1 {
				hclValid = true
			}
		}

		if passport.Ecl != "" {
			re, err := regexp.Compile("^amb$|^blu$|^brn$|^gry$|^grn$|^hzl$|^oth$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Ecl)
			if len(matches) == 1 {
				eclValid = true
			}
		}

		if passport.Pid != "" {
			re, err := regexp.Compile("^\\d{9}$")
			if err != nil {
				log.Fatal(err)
			}
			matches := re.FindStringSubmatch(passport.Pid)
			if len(matches) == 1 {
				pidValid = true
			}
		}

		if byrValid && iyrValid && eyrValid && hgtValid && hclValid && eclValid && pidValid {
			answer++
		}
	}

	return answer
}

func doWork(passportInfo string) passport {
	var pid, hcl, byr, iyr, eyr, hgt, ecl, cid string
	parts := strings.Split(passportInfo, " ")
	for _, part := range parts {
		if strings.Contains(part, "pid") {
			pid = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "hcl") {
			hcl = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "byr") {
			byr = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "iyr") {
			iyr = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "hgt") {
			hgt = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "ecl") {
			ecl = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "eyr") {
			eyr = strings.Split(part, ":")[1]
		} else if strings.Contains(part, "cid") {
			cid = strings.Split(part, ":")[1]
		}
	}
	return passport{
		Byr: byr,
		Iyr: iyr,
		Eyr: eyr,
		Hgt: hgt,
		Hcl: hcl,
		Ecl: ecl,
		Pid: pid,
		Cid: cid,
	}
}

func day3() {
	answer3 := crawlRightByX(3, false)
	answer1 := crawlRightByX(1, false)
	answer5 := crawlRightByX(5, false)
	answer7 := crawlRightByX(7, false)
	answer12 := crawlRightByX(1, true)

	fmt.Printf("day3 part 1 answer: %d\n", answer3)
	fmt.Printf("day3 part 2 answer: %d*%d*%d*%d*%d=%d\n", answer1, answer5, answer3, answer7, answer12,
		answer1*answer5*answer3*answer7*answer12)
}

func crawlRightByX(crawl int, twoDown bool) int {
	f, err := os.Open("day3.txt")
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

func day2() {
	type password struct {
		Min    int    `json:",min"`
		Max    int    `json:",max"`
		Letter string `json:",letter"`
		Value  string `json:",value"`
	}

	f, err := os.Open("day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var passwords []password
	for scanner.Scan() {
		s := scanner.Text()
		parts := strings.Split(s, " ")

		minMax := strings.Split(parts[0], "-")
		letter := string(parts[1][0])

		min, err := strconv.Atoi(minMax[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			log.Fatal(err)
		}

		passwords = append(passwords, password{
			Min:    min,
			Max:    max,
			Letter: letter,
			Value:  parts[2],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var answer int
	for _, p := range passwords {
		var numFound int
		for _, s := range p.Value {
			if p.Letter == string(s) {
				numFound++
			}
		}

		if numFound >= p.Min && numFound <= p.Max {
			answer++
		}
	}

	fmt.Printf("day2 part 1 answer: %d\n", answer)

	var answer2 int
	for _, p := range passwords {
		char1 := string(p.Value[p.Min-1])
		char2 := string(p.Value[p.Max-1])

		if (p.Letter == char1 && p.Letter != char2) || (p.Letter != char1 && p.Letter == char2) {
			answer2++
			continue
		}
	}

	fmt.Printf("day2 part 2 answer: %d\n", answer2)
}

func day1() {
	f, err := os.Open("day1.txt")
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
