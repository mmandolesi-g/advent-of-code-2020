package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Node struct {
	Name     string
	Children map[*Node]int
	Parents  []*Node
}

func main() {
	f, err := os.Open("./day-07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	re := regexp.MustCompile(`((\d*\s*)(\w+\s\w+\s)bag(s)*)`)
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
