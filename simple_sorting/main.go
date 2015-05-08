package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(sorted(s))
	}
}

func sorted(s string) string {
	parts := strings.Fields(s)
	sort.Sort(MySlice(parts))
	return strings.Join(parts, " ")
}

type MySlice []string

func (p MySlice) Len() int { return len(p) }
func (p MySlice) Less(i, j int) bool {
	x, _ := strconv.ParseFloat(p[i], 64)
	y, _ := strconv.ParseFloat(p[j], 64)
	return x < y
}
func (p MySlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }
