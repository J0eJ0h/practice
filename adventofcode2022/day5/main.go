package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	f, err := os.Open(".\\input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %s\n", err)
		return nil
	}
	defer f.Close()

	var input []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		input = append(input, s.Text())
	}

	return input
}

type BoxQueue struct {
	boxen []rune
}

func (q *BoxQueue) Peek() rune {
	return q.boxen[len(q.boxen)-1]
}

func (q *BoxQueue) Pop() rune {
	r := q.Peek()
	q.boxen = q.boxen[:len(q.boxen)-1]
	return r
}

func (q *BoxQueue) Push(r rune) {
	q.boxen = append(q.boxen, r)
}

func splitInput(input []string) (boxLayout, moves []string) {
	for i, v := range input {
		if v == "" {
			return input[:i], input[i+1:]
		}
	}
	return nil, nil
}

func parseBoxen(bs []string) []BoxQueue {
	// so bad
	boxen := make([]BoxQueue, 9) // cry

	for i := len(bs) - 2; i >= 0; i-- {
		for j := 0; j < 9; j++ {
			c := rune(bs[i][4*j+1])
			if c != ' ' {
				boxen[j].Push(c)
			}
		}
	}
	return boxen
}

type Move struct {
	count, from, to int
}

func parseMoves(ms []string) (moves []Move) {
	for _, s := range ms {
		move := strings.Fields(s)
		count, _ := strconv.Atoi(move[1])
		from, _ := strconv.Atoi(move[3])
		to, _ := strconv.Atoi(move[5])
		moves = append(moves, Move{count: count, from: from, to: to})
	}
	return
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	// TODO: write actual parsers for each of the 2 input stages with grammar instead of lexing it
	// Never mind. This format is hard B$.
	a, b := splitInput(input)
	for i, v := range a {
		fmt.Printf("%v %v\n", i, len(v))
	}
	boxen := parseBoxen(a)
	fmt.Printf("Boxen: %v\n", boxen)

	moves := parseMoves(b)
	for _, v := range moves {
		for i := 0; i < v.count; i++ {
			boxen[v.to-1].Push(boxen[v.from-1].Pop())
		}
	}
	result := ""
	for _, v := range boxen {
		result += string(v.Peek())
	}
	fmt.Printf("Result: %v\n", result)
}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))

	a, b := splitInput(input)
	for i, v := range a {
		fmt.Printf("%v %v\n", i, len(v))
	}
	boxen := parseBoxen(a)
	fmt.Printf("Boxen: %v\n", boxen)

	moves := parseMoves(b)
	bq := BoxQueue{}
	for _, v := range moves {
		for i := 0; i < v.count; i++ {
			bq.Push(boxen[v.from-1].Pop())
		}
		for i := 0; i < v.count; i++ {
			boxen[v.to-1].Push(bq.Pop())
		}
	}
	result := ""
	for _, v := range boxen {
		result += string(v.Peek())
	}
	fmt.Printf("Result: %v\n", result)
}

func main() {
	dopart := flag.Int("part", 2, "Specify question part")

	switch *dopart {
	case 1:
		part1()
	case 2:
		part2()
	default:
		fmt.Println("invalid part")
	}
}
