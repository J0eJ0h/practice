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

func atoiArr(sa []string) (ret []int) {
	for _, v := range sa {
		n, _ := strconv.Atoi(v)
		ret = append(ret, n)
	}
	return
}

func parseLine(s string) (a, b, x, y int) {
	e := atoiArr(strings.FieldsFunc(s, func(r rune) bool { return (r == '-' || r == ',') }))
	a, b, x, y = e[0], e[1], e[2], e[3]
	return
}

func fullycontains(a, b, x, y int) bool {
	return a <= x && b >= y || x <= a && y >= b
}

func overlap(a, b, x, y int) bool {
	if a > b || x > y {
		log.Fatal("WAT")
	}
	return a >= x && a <= y || b >= x && b <= y
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	botaltot := 0
	for _, v := range input {
		a, b, x, y := parseLine(v)
		if fullycontains(a, b, x, y) {
			botaltot++
		}
	}

	fmt.Printf("Botaltot: %v\n", botaltot)

}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))

	botaltot := 0
	for _, v := range input {
		a, b, x, y := parseLine(v)
		if fullycontains(a, b, x, y) || overlap(a, b, x, y) {
			botaltot++
		}
	}

	fmt.Printf("Botaltot: %v\n", botaltot)
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
