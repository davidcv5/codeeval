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
		s := scanner.Text()
		fmt.Println(distance(s))
	}
}

func distance(s string) int {
	s1 := strings.Replace(s, "(", "", -1)
	s1 = strings.Replace(s1, ")", "", -1)
	s1 = strings.Replace(s1, ",", "", -1)
	parts := strings.Fields(s1)
	x1, _ := strconv.Atoi(parts[0])
	y1, _ := strconv.Atoi(parts[1])
	x2, _ := strconv.Atoi(parts[2])
	y2, _ := strconv.Atoi(parts[3])

	return int(math.Sqrt(math.Pow(math.Abs(float64(x1-x2)), 2) + math.Pow(math.Abs(float64(y1-y2)), 2)))
}
