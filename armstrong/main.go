package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
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
		fmt.Println(armstrong(s))
	}
}

func armstrong(s string) string {
	x, _ := strconv.Atoi(s)
	n := len(s)

	for _, r := range s {
		x -= int(math.Pow(float64(r-'0'), float64(n)))
	}

	if x == 0 {
		return "True"
	}
	return "False"
}
