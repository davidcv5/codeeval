package main

import (
	"fmt"
	"math"
	"strconv"
)
import "log"
import "bufio"
import "os"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	cases := 0
	if scanner.Scan() {
		cases, _ = strconv.Atoi(scanner.Text())
	}
	for scanner.Scan() {
		if cases == 0 {
			break
		}
		fmt.Println(doubleSquares(scanner.Text()))
		cases--
	}
}

func doubleSquares(in string) int {
	pairs := 0
	x, _ := strconv.ParseFloat(in, 64)
	max := int(math.Sqrt(x))
	m := make(map[float64]int)
	for i := max; i >= 0; i-- {
		p := math.Pow(float64(i), 2)
		m[p] = i
		if _, ok := m[x-p]; ok {
			pairs++
		}

	}
	return pairs
}
