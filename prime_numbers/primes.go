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
		fmt.Println(primes(scanner.Text()))
	}
}

func primes(s string) string {
	n, _ := strconv.Atoi(s)
	p := []string{}
	for i := 2; i < n; i++ {
		if isPrime(i) {
			p = append(p, strconv.Itoa(i))
		}
	}
	return strings.Join(p, ",")
}

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	i := 5
	w := 2
	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i += 2
		w = 6 - w
	}
	return true
}
