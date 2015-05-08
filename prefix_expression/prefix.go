package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var operators = "+*/"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := scanner.Text()
		fmt.Println(evaluate(s))
	}
}

func evaluate(s string) int {
	parts := strings.Fields(s)
	ops := []string{}
	nums := []int{}

	for _, p := range parts {
		if strings.Contains(operators, p) {
			ops = append(ops, p)
		} else {
			n, _ := strconv.Atoi(p)
			nums = append(nums, n)
		}
	}

	result := nums[0]
	nums = nums[1:]
	for i := 0; i < len(nums); i++ {
		o := ops[i]
		n := nums[i]
		for ; i < len(nums)-1 && order(o) < order(ops[i+1]); i++ {
			n = eval(n, nums[i+1], ops[i+1])
		}
		result = eval(result, n, o)
	}
	return result
}

func eval(x, y int, op string) int {
	switch op {
	case "+":
		return x + y
	case "*":
		return x * y
	case "/":
		if y == 0 {
			return 0
		}
		return x / y
	}
	return 0
}

func order(o string) int {
	if o == "+" {
		return 0
	}
	return 1
}
