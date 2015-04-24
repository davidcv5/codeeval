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
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		fmt.Println(multiples(line))
	}
}

func multiples(args string) int {
	a := strings.Split(args, ",")
	var (
		err  error
		x, y int
	)
	x, err = strconv.Atoi(a[0])
	y, err = strconv.Atoi(a[1])

	if err != nil {
		log.Fatal(err)
	}

	for i := 1; i < x; i++ {
		if x <= y*i {
			fmt.Println(i)
			return y * i
		}
	}

	return 0
}
