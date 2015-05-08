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
		fmt.Println(isSelfDesc(scanner.Text()))
	}
}

func isSelfDesc(in string) int {
	m := make(map[int]int)
	for i, c := range in {
		v := int(c - '0')

		if _, ok := m[i]; ok {
			m[i] += v
		} else {
			m[i] = v
		}

		if _, ok := m[v]; ok {
			m[v] -= 1
		} else {
			m[v] = -1
		}
	}

	for _, v := range m {
		if v != 0 {
			return 0
		}
	}
	return 1
}
