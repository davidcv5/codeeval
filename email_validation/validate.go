package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
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
		fmt.Println(validate(s))
	}
}

func validate(s string) bool {
	r := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+\@[a-zA-Z0-9-.]+\.[a-zA-Z0-9]+$`)
	return r.MatchString(s)
}
