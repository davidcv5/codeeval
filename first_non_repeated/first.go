package main

import "fmt"
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
		fmt.Println(first(scanner.Text()))
	}
}

func first(in string) string {
	m := make(map[rune]int)
	for _, v := range in {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	for _, v := range in {
		if m[v] == 1 {
			return string(v)
		}
	}
	return ""
}
