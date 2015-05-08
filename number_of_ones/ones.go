package main

import (
	"fmt"
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
	for scanner.Scan() {
		//'scanner.Text()' represents the test case, do something with it
		fmt.Println(ones(scanner.Text()))
	}
}

func ones(in string) int {
	count := 0
	x, _ := strconv.ParseInt(in, 10, 64)
	s := strconv.FormatInt(x, 2)
	for _, c := range s {
		if c == '1' {
			count++
		}
	}
	return count
}
