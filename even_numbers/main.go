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
		s := scanner.Text()
		x, _ := strconv.Atoi(s)
		if x%2 == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
