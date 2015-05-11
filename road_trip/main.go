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
		fmt.Println(trip(scanner.Text()))
	}
}

func trip(s string) string {
	parts := strings.Split(s, ";")
	d := []int{}
	for _, c := range parts[:len(parts)-1] {
		p := c[strings.Index(c, ",")+1:]
		x, _ := strconv.Atoi(p)
		d = append(d, x)
	}
	sort.Ints(d)
	result := make([]string, len(d))
	x := d[0]
	result[0] = strconv.Itoa(x)
	for i := 1; i < len(d); i++ {
		result[i] = strconv.Itoa(d[i] - x)
		x = d[i]
	}
	return strings.Join(result, ",")
}
