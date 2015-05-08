package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)
import "bufio"
import "os"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make(bySize, 0)
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	sort.Sort(lines)
	for i := 0; i < n; i++ {
		fmt.Println(lines[i])
	}
}

type bySize []string

func (a bySize) Len() int           { return len(a) }
func (a bySize) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a bySize) Less(i, j int) bool { return len(a[i]) > len(a[j]) }
