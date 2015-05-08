package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s, c := reverse_and_add(scanner.Text(), 0)
		fmt.Println(s, c)
	}
}

func reverse_and_add(s string, step int) (int, string) {
	if palindrome(s) {
		return step, s
	}
	x, _ := strconv.Atoi(s)
	y, _ := strconv.Atoi(reverse(s))

	return reverse_and_add(strconv.Itoa(x+y), step+1)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func palindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < len(runes)/2; {
		if runes[i] != runes[j] {
			return false
		}
		i++
		j--
	}
	return true
}
