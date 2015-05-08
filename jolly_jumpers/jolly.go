package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
		fmt.Println(jolly(scanner.Text()))
	}
}

func jolly(s string) string {
	parts := strings.Fields(s)
	n, _ := strconv.Atoi(parts[0])
	if n == 1 {
		return "Jolly"
	}

	parts = parts[1:]
	total := (n) * (n - 1) / 2
	for i := 0; i < len(parts)-1; i++ {
		x, _ := strconv.Atoi(parts[i])
		y, _ := strconv.Atoi(parts[i+1])
		total -= int(math.Abs(float64(y - x)))
	}
	if total == 0 {
		return "Jolly"
	}
	return "Not jolly"
}
