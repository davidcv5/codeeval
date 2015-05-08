package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
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
		fmt.Println(h2d(s))
	}
}

func h2d(s string) int {
	r := 0
	for i := len(s) - 1; i >= 0; i-- {
		r += getDec(s[i]) * int(math.Pow(float64(16), float64(len(s)-1-i)))
	}
	return r
}

func getDec(r uint8) int {
	switch r {
	case 'a':
		return 10
	case 'b':
		return 11
	case 'c':
		return 12
	case 'd':
		return 13
	case 'e':
		return 14
	case 'f':
		return 15
	}
	return int(r - '0')
}
