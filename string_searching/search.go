package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		x := s[:strings.Index(s, ",")]
		y := s[strings.Index(s, ",")+1:]

		fmt.Println(search(x, y))
	}
}

func search(x, y string) bool {
	for i := 0; i < len(x); i++ {
		k := i
		j := 0
		for ; j < len(y) && i < len(x); j++ {
			if y[j] == '*' {
				if j == len(y)-1 {
					return true
				}
				return search(x[i+1:], y[j+1:])
				//fmt.Println(string(x[i]), string(y[j]))
			} else if y[j] == '\\' && j < len(y)-1 && y[j+1] == '*' && x[i] == '*' {
				//fmt.Println("escaping", string(x[i]), string(y[j]))
				i++
				j++
			} else if x[i] == y[j] {
				//fmt.Println(string(x[i]), string(y[j]))
				i++
			} else {
				break
			}
		}
		if j == len(y) {
			return true
		}
		i = k
	}
	return false
}
