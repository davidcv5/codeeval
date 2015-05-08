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
		fmt.Println(unique(scanner.Text()))
	}
}

func unique(in string) string {
	m := make(map[string]bool)
	result := ""
	for _, k := range strings.Split(in, ",") {
		if _, ok := m[k]; !ok {
			result += k + ","
			m[k] = true
		}
	}
	return strings.Trim(result, ",")
}
