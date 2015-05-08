package main

import (
	"fmt"
	"strconv"
	"strings"
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
	for scanner.Scan() {
		//'scanner.Text()' represents the test case, do something with it
		fmt.Println(sum(scanner.Text()))
	}
}

func sum(in string) int {
	parts := strings.Split(in, ",")
	max, _ := strconv.Atoi(parts[0])
	current := max
	for i := 1; i < len(parts); i++ {
		x, _ := strconv.Atoi(parts[i])
		if current < 0 && current < x {
			current = x
		} else {
			current += x
		}
		if current > max {
			max = current
		}
	}
	return max
}
