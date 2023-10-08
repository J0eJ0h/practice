package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
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

type WindowBuffer struct {
	buff    []rune
	maxSize int
}

func (wb *WindowBuffer) Add(r rune) {
	wb.buff = append(wb.buff, r)
	if len(wb.buff) > wb.maxSize {
		wb.buff = wb.buff[1:]
	}
}

func (wb *WindowBuffer) FilledUniq() bool {
	if len(wb.buff) != wb.maxSize {
		return false
	}
	for i := 0; i < wb.maxSize-1; i++ {
		for j := i + 1; j < wb.maxSize; j++ {
			if wb.buff[i] == wb.buff[j] {
				return false
			}
		}
	}
	return true
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	s := input[0]
	wb := WindowBuffer{maxSize: 4}
	for i, v := range s {
		wb.Add(v)
		if wb.FilledUniq() {
			fmt.Printf("Found char %v\n", i+1)
		}
	}
}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))

	s := input[0]
	wb := WindowBuffer{maxSize: 14}
	for i, v := range s {
		wb.Add(v)
		if wb.FilledUniq() {
			fmt.Printf("Found char %v\n", i+1)
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
