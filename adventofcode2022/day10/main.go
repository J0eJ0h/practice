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

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error paring %v\n", err)
	}
	return i
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	X := 1
	cycle := 0
	tots := 0
	for _, v := range input {
		f := strings.Fields(v)
		cycle++
		if (cycle+20)%40 == 0 {
			tots += cycle * X
		}
		switch f[0] {
		case "addx":
			cycle++
			if (cycle+20)%40 == 0 {
				tots += cycle * X
			}
			X += atoi(f[1])
		}
	}

	fmt.Printf("Answwer %v\n", tots)

}

func update(cycle, X int) {
	column := (cycle - 1) % 40

	if column == X-1 || column == X || column == X+1 {
		fmt.Print("X")
	} else {
		fmt.Print(".")
	}
	if column == 39 {
		fmt.Println()
	}
}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))

	X := 1
	cycle := 0
	for _, v := range input {
		f := strings.Fields(v)
		cycle++
		update(cycle, X)
		switch f[0] {
		case "addx":
			cycle++
			update(cycle, X)
			X += atoi(f[1])
		}
	}
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
