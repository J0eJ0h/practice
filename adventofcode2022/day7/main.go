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

type File struct {
	size int
}

func (f *File) Size() int {
	return f.size
}

type Directory struct {
	name       string
	childrenz  map[string]*Directory
	parent     *Directory
	files      map[string]*File
	cachedSize int
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{name: name, childrenz: make(map[string]*Directory), files: make(map[string]*File), parent: parent}
}

func (d *Directory) AddChildDir(name string) (child *Directory) {
	var ok bool
	if child, ok = d.childrenz[name]; !ok {
		child = NewDirectory(name, d)
		d.childrenz[name] = child
	}
	return child
}

func (d *Directory) AddFile(name string, size int) {
	d.files[name] = &File{size: size}
}

func (d *Directory) GetPath() string {
	if d.parent == nil {
		return ""
	}
	return d.parent.GetPath() + "/" + d.name
}

func (d *Directory) Size(useCache bool) int {
	if useCache {
		if d.parent != nil {
			fmt.Printf("%v:%v\n", d.GetPath(), d.cachedSize)
		}
		return d.cachedSize
	}
	size := 0
	for _, f := range d.files {
		size += f.Size()
	}
	for _, c := range d.childrenz {
		size += c.Size(false)
	}
	d.cachedSize = size
	return size
}

func buildFS(input []string) (root *Directory) {
	root = NewDirectory("/", nil)
	cwd := root
	for _, l := range input {
		toke := strings.Fields(l)
		if toke[0] != "$" {
			if toke[0] == "dir" {
				cwd.AddChildDir(toke[1])
			} else {
				size, _ := strconv.Atoi(toke[0])
				cwd.AddFile(toke[1], size)
			}
			continue
		}
		if toke[1] == "cd" {
			if toke[2] == ".." {
				cwd = cwd.parent
			} else if toke[2] == "/" {
				cwd = root
			} else {
				cwd = cwd.AddChildDir(toke[2])
			}
		}
	}
	return
}

func (d *Directory) CountSize(threshold int) int {
	tots := 0
	if d.Size(true) <= threshold {
		tots += d.Size(true)
	}
	for _, v := range d.childrenz {
		tots += v.CountSize(threshold)
	}
	return tots
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	root := buildFS(input)
	fmt.Printf("Size: %v\n", root.Size(false))

	fmt.Printf("Answer: %v\n", root.CountSize(100000))

}

func (d *Directory) FindTwack(toFree int) (string, int) {
	name, size := d.name, d.Size(true)
	if size < toFree {
		return "", 0
	}

	for _, v := range d.childrenz {
		tmpn, tmpv := v.FindTwack(toFree)
		if tmpn == "" {
			continue
		}
		if tmpv < size {
			name, size = tmpn, tmpv
		}
	}
	return name, size
}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))

	root := buildFS(input)
	fmt.Printf("Size: %v\n", root.Size(false))

	rootSize := root.Size(false)
	freeSpace := 70000000 - rootSize
	toFree := 30000000 - freeSpace
	fmt.Printf("Size: %v    Free Space: %v    toFree: %v\n", rootSize, freeSpace, toFree)

	name, size := root.FindTwack(toFree)
	fmt.Printf("Directory to delete: %v %v out of %v\n", name, size, toFree)

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
