package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var operators = []string{"+", "-", ""}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(ugly(s))
	}
}

func ugly(s string) int {
	count := 0
	options(s[:1], s[1:], &count)
	//fmt.Println(count)
	return count
}

func options(pref, suf string, count *int) {

	if len(suf) == 0 {
		if isUgly(pref) {
			*count += 1
		}
		return
	}
	for _, op := range operators {
		options(pref+op+suf[:1], suf[1:], count)
	}
}

func isUgly(s string) bool {
	n := eval(s)
	ugly := false
	if n%2 == 0 || n%3 == 0 || n%5 == 0 || n%7 == 0 {
		ugly = true
	}
	//fmt.Println("isUgly", s, n, ugly)

	return ugly
}

func eval(s string) int {
	sum := 0
	n := ""
	op := '+'
	for _, r := range s {
		if r != '+' && r != '-' {
			n += string(r)
		} else {
			sum = calculate(sum, n, op)
			op = r
			n = ""
		}
	}
	return calculate(sum, n, op)
}

func calculate(sum int, n string, op rune) int {
	x, _ := strconv.Atoi(n)
	switch op {
	case '+':
		//fmt.Println(sum, string(op), n, "=", sum+x)
		return sum + x
	case '-':
		//fmt.Println(sum, string(op), n, "=", sum-x)
		return sum - x
	}
	return 0
}
