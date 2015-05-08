package main

import (
	"bufio"
	"fmt"
	"log"
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
		line := scanner.Text()
		if len(line) != 0 {
			fmt.Println(pairs(scanner.Text()))
		}
	}
}

func pairs(s string) string {
	parts := strings.Split(s, ";")
	x, _ := strconv.Atoi(parts[1])
	numbers := strings.Split(parts[0], ",")
	result := []string{}
	for i, j := 0, len(numbers)-1; i < j; {
		a, _ := strconv.Atoi(numbers[i])
		b, _ := strconv.Atoi(numbers[j])
		if a+b > x {
			j--
		} else if a+b < x {
			i++
		} else {
			result = append(result, fmt.Sprintf("%d,%d", a, b))
			i++
			j--
		}
	}
	if len(result) > 0 {
		return strings.Join(result, ";")
	}
	return "NULL"
}
