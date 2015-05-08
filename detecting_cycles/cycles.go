package main

import (
	"fmt"
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
		fmt.Println(getCycle(scanner.Text()))
	}
}

func getCycle(line string) string {
	s := strings.Split(line, " ")

	m := make(map[int]int)

	for i, v := range s {

	}
	return ""
}
