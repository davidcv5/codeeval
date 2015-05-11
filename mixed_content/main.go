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
		fmt.Println(mixed(scanner.Text()))
	}
}

func mixed(s string) string {
	parts := strings.Split(s, ",")
	c := []string{}
	d := []string{}
	for _, v := range parts {
		if _, err := strconv.Atoi(v); err == nil {
			d = append(d, v)
		} else {
			c = append(c, v)
		}
	}
	if len(c) == 0 || len(d) == 0 {
		return s
	}
	return strings.Join(c, ",") + "|" + strings.Join(d, ",")
}
