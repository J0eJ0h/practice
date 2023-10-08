package main

import (
	"flag"
	"fmt"
)

func part1() {
	fmt.Println("part1")
}

func part2() {
	fmt.Println("part2")
}

func main() {
	dopart := flag.Int("part", 1, "Specify question part")

	switch *dopart {
	case 1:
		part1()
	case 2:
		part2()
	default:
		fmt.Println("invalid part")
	}
}
