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
		fmt.Println(intersection(scanner.Text()))
	}
}

func intersection(in string) string {
	parts := strings.Split(in, ";")
	a := strings.Split(parts[0], ",")
	b := strings.Split(parts[1], ",")

	result := ""
	for i, j := 0, 0; i < len(a) && j < len(b); {
		x, err := strconv.Atoi(a[i])
		y, err := strconv.Atoi(b[j])
		if err != nil {
			return ""
		}
		if x == y {
			result += fmt.Sprintf("%d,", x)
			i++
			j++
		} else if x < y {
			i++
		} else {
			j++
		}
	}
	return strings.Trim(result, ",")
}
