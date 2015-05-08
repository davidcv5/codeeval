package main

import (
	"fmt"
	"strconv"
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
		if i, err := strconv.Atoi(scanner.Text()); err == nil {
			fmt.Println(isHappy(i))
		}
	}
}

func isHappy(i int) int {
	m := make(map[int]bool)

	x := i
	for {
		if x == 1 {
			return 1
		}
		if _, ok := m[x]; ok {
			return 0
		}
		m[x] = true
		s := strconv.Itoa(x)
		x = 0
		for _, v := range s {
			n := int(v - '0')
			x += n * n
		}
	}
}
