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
		s := scanner.Text()
		if len(s) > 0 {
			fmt.Println(absurdity(s))
		}
	}
}

func absurdity(s string) int {
	parts := strings.Split(s, ";")
	size, _ := strconv.Atoi(parts[0])
	sum := 0
	for _, r := range strings.Split(parts[1], ",") {
		n, _ := strconv.Atoi(r)
		sum += n
	}
	return sum - ((size-1)*(size-2))/2
}
