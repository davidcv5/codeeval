package main

import (
	"fmt"
	"strconv"
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
		line := strings.Fields(scanner.Text())
		n, err := strconv.Atoi(line[len(line)-1])
		line = line[:len(line)-1]
		if err != nil || n > len(line) {
			continue
		}
		fmt.Println(line[len(line)-n])
	}
}
