package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

func part1() []int {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	var elves []int
	currentElf := 0
	maxElf := 0
	for _, s := range input {
		if s == "" {
			elves = append(elves, currentElf)
			if currentElf > maxElf {
				maxElf = currentElf
			}
			currentElf = 0
		} else {
			doof, _ := strconv.Atoi(s)
			currentElf += doof
		}

	}
	if currentElf != 0 {
		elves = append(elves, currentElf)
		if currentElf > maxElf {
			maxElf = currentElf
		}
	}
	fmt.Printf("Max Elf: %v\n", maxElf)
	return elves
}

func part2() {
	fmt.Println("part2")

	elves := part1()
	sort.Ints(elves)
	fmt.Printf("Ha: %v\n", elves[len(elves)-3:])
	sum := 0
	for _, v := range elves[len(elves)-3:] {
		sum += v
	}
	fmt.Printf("Doof %v\n", sum)

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
