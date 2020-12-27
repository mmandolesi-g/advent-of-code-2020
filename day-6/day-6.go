package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("./day-6/input.txt")
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
	day6Part2()
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

func day6Part2() {
	f, err := os.Open("./day-6/input.txt")
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
