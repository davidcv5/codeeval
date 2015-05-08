package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		fmt.Println(next(s))
	}
}

func next(s string) string {
	if len(s) == 1 {
		return s + "0"
	}
	permutations := []int{}
	permutate("", s, &permutations)
	sort.Ints(permutations)
	i := 0
	x, _ := strconv.Atoi(s)
	for ; i < len(permutations) && permutations[i] <= x; i++ {
	}
	if i == len(permutations) {
		return s[:1] + "0" + s[1:]
	}
	return strconv.Itoa(permutations[i])
}

func permutate(pref, suf string, permutations *[]int) {
	if len(suf) == 0 {
		if pref[0] != '0' {
			x, _ := strconv.Atoi(pref)
			*permutations = append(*permutations, x)
		}
		return
	}
	for i, r := range suf {
		permutate(pref+string(r), suf[:i]+suf[i+1:], permutations)
	}
}
