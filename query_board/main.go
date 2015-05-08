package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var board = [256 * 256]int{}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		query(s)
	}
}

func query(s string) {
	parts := strings.Fields(s)
	i, _ := strconv.Atoi(parts[1])
	x := 0
	if len(parts) > 2 {
		x, _ = strconv.Atoi(parts[2])
	}

	switch parts[0] {
	case "SetCol":
		setCol(i, x)
		break
	case "SetRow":
		setRow(i, x)
		break
	case "QueryCol":
		queryCol(i)
		break
	case "QueryRow":
		queryRow(i)
		break
	}
}

func setCol(i, x int) {
	for j := 0; j < 256; j++ {
		board[i+256*j] = x
	}
}

func setRow(i, x int) {
	for j := 0; j < 256; j++ {
		board[i*256+j] = x
	}
}

func queryCol(i int) {
	result := 0
	for j := 0; j < 256; j++ {
		result += board[i+256*j]
	}
	fmt.Println(result)
}

func queryRow(i int) {
	result := 0
	for j := 0; j < 256; j++ {
		result += board[i*256+j]
	}
	fmt.Println(result)
}
