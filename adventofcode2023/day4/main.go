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
	f, err := os.Open("./input.txt")
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

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error paring %v\n", err)
	}
	return i
}

func digit(r rune) int {
	i, err := strconv.Atoi(string(r))
	if err != nil {
		return -1
	}
	return i
}

func symbol(r rune) bool {
	if digit(r) > -1 || r == '.' {
		return false
	}
	return true
}

type Engine struct {
	w, h int
	e    [][]rune
}

func getEngine(input []string) *Engine {
	e := make([][]rune, len(input))
	for i, l := range input {
		e[i] = []rune(strings.TrimSpace(l))
	}
	return &Engine{w: len(e[0]), h: len(e), e: e}
}

func (e *Engine) hasAdjSymbol(row, start, end int) (bool, rune, int, int) {
	if start > 0 {
		start -= 1
	}

	// end is inclusive
	if end < e.w-1 {
		end += 1
	}
	// Don't check the provided elements
	// fmt.Println(end)
	r := e.e[row][start]
	if symbol(r) {
		return true, r, row, start
	}
	r = e.e[row][end]
	if symbol(r) {
		return true, r, row, end
	}

	if row > 0 {
		for i := start; i <= end; i++ {
			r = e.e[row-1][i]
			if symbol(r) {
				return true, r, row - 1, i
			}
		}
	}
	if row < e.h-1 {
		for i := start; i <= end; i++ {
			r = e.e[row+1][i]
			if symbol(r) {
				return true, r, row + 1, i
			}
		}
	}
	return false, '.', 0, 0
}

func (e *Engine) getNumber(row, start int) (int, int) {
	total, i := 0, start
	//fmt.Println(string(e.e[row][start:]))

	for i = start; i < e.w; i++ {
		d := digit(e.e[row][i])
		//fmt.Printf("%d %c %d\n", i, e.e[row][i], d)
		if d < 0 {
			break
		}
		total *= 10
		total += d
	}

	if total > 0 {
		i -= 1
		fmt.Println(i)
	}
	return total, i
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))
	dopart1(input)
}
func dopart1(input []string) {
	total := 0

	for _, l := range input {
		total += readCard(l).scoreCard()
	}

	fmt.Println(total)
}

type Card struct {
	winners map[string]struct{}
	have    []string
	val     int
}

func readCard(l string) *Card {
	id_rest := strings.Split(l, ":")
	win_have := strings.Split(id_rest[1], "|")
	winners, have := strings.Fields(win_have[0]), strings.Fields(win_have[1])

	card := Card{winners: make(map[string]struct{}), have: have, val: 1}
	for _, v := range winners {
		card.winners[v] = struct{}{}
	}
	return &card
}

func (c *Card) countCard() int {
	total := 0
	for _, v := range c.have {
		if _, ok := c.winners[v]; ok {
			total++
		}
	}
	return total
}

func (c *Card) scoreCard() int {
	total := c.countCard()
	if total > 0 {
		return 1 << (total - 1)
	} else {
		return 0
	}
}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))

	total := 0
	cards := make([]*Card, len(input))
	for i, v := range input {
		cards[i] = readCard(v)
	}

	for i := len(cards) - 1; i >= 0; i-- {
		for j := 0; j < cards[i].countCard(); j++ {
			cards[i].val += cards[i+j+1].val
		}
		total += cards[i].val
	}

	fmt.Println(total)
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

	/*test :=
		`.........
	 123@.456$
	 .........` */

	/*test = `467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..` */

	//dopart1(strings.Split(test, "\n"))
}
