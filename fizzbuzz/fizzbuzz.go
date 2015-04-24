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
		line := strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		fmt.Println(fizzBuzz(line))
	}
}

func fizzBuzz(args string) string {
	a := strings.Split(args, " ")
	var (
		err     error
		x, y, z int
	)
	x, err = strconv.Atoi(a[0])
	y, err = strconv.Atoi(a[1])
	z, err = strconv.Atoi(a[2])

	if err != nil {
		log.Fatal(err)
	}

	result := make([]string, z)
	for i := 1; i <= z; i++ {

		if i%x == 0 && i%y == 0 {
			result[i-1] = "FB"
		} else if i%x == 0 {
			result[i-1] = "F"
		} else if i%y == 0 {
			result[i-1] = "B"
		} else {
			result[i-1] = strconv.Itoa(i)
		}
	}
	return strings.Join(result, " ")
}
