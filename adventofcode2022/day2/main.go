package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

func score(them, me int) int {
	switch them - me {
	case -2, 1:
		return me
	case 0:
		return me + 3
	case -1, 2:
		return me + 6
	}

	log.Fatalf("Bad code %v %v \n", them, me)
	return 0
}

func parseChar(c string) int {
	switch c {
	case "A", "X":
		return 1
	case "B", "Y":
		return 2
	case "C", "Z":
		return 3
	}
	return 42
}

func scoreLine(l string) int {
	f := strings.Fields(l)
	return score(parseChar(f[0]), parseChar(f[1]))
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	total := 0
	for _, v := range input {
		total += scoreLine(v)
	}

	fmt.Printf("Score: %v\n", total)

}

func scoreLine2(l string) int {
	f := strings.Fields(l)
	them := parseChar(f[0])
	me := 42
	switch f[1] {
	case "X": // lose
		me = them + 2
	case "Y":
		me = them
	case "Z":
		me = them + 1
	default:
		log.Fatal("WAT")
	}
	return score(them, ((me-1)%3)+1)
}
func part2() {
	fmt.Println("part2")

	input := readInput()
	total := 0
	for _, v := range input {
		total += scoreLine2(v)
	}
	fmt.Printf("Score %v\n", total)
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
