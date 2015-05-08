package main

import (
	"fmt"
	"strconv"
	"strings"
)
import "log"
import "bufio"
import "os"

type btree struct {
	left, right *btree
	value       int
}

var tree = &btree{
	value: 30,
	left: &btree{
		value: 8,
		left: &btree{
			value: 3,
		},
		right: &btree{
			value: 20,
			left: &btree{
				value: 10,
			},
			right: &btree{
				value: 29,
			},
		},
	},
	right: &btree{
		value: 52,
	},
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//'scanner.Text()' represents the test case, do something with it
		a := make([]string, 0)
		b := make([]string, 0)
		parts := strings.Fields(scanner.Text())
		x, _ := strconv.Atoi(parts[0])
		y, _ := strconv.Atoi(parts[1])
		ancestors(tree, x, &a)
		ancestors(tree, y, &b)
		result := ""
		for i, j := len(a)-1, len(b)-1; i >= 0 && j >= 0 && a[i] == b[j]; {
			result = a[i]
			i--
			j--
		}
		fmt.Println(result)
	}
}

func ancestors(b *btree, x int, a *[]string) bool {
	if b.value == x {
		*a = append(*a, strconv.Itoa(x))
		return true
	}
	if b.left != nil && ancestors(b.left, x, a) {
		*a = append(*a, strconv.Itoa(b.value))
		return true
	}
	if b.right != nil && ancestors(b.right, x, a) {
		*a = append(*a, strconv.Itoa(b.value))
		return true
	}
	return false
}
