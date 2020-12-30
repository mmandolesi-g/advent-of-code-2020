package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("./day-13/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var (
		buses        []bus
		i, startTime int
	)
	for scanner.Scan() {
		line := scanner.Text()

		if i == 0 {
			startTime, err = strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			i++
			continue
		}

		schedules := strings.Split(line, ",")
		for _, s := range schedules {
			if s == "x" {
				continue
			}

			ss, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}

			for x := 0; x < startTime+ss; x = x + ss {
				if x < startTime {
					continue
				}
				buses = append(buses, bus{
					id:   ss,
					time: x - startTime,
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(buses[:], func(i, j int) bool {
		return buses[i].time < buses[j].time
	})

	log.Printf("day13 part 1 answer: %d", buses[0].time*buses[0].id)
}

type bus struct {
	id   int
	time int
}
