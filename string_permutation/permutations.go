package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
		s := scanner.Text()
		if len(s) > 0 {
			fmt.Println(permutations(s))
		}
	}
}

func permutations(s string) string {
	p := []string{}
	permutate("", s, &p)
	sort.Strings(p)
	return strings.Join(p, ",")
}

func permutate(pref, suf string, p *[]string) {
	if len(suf) == 0 {
		*p = append(*p, pref)
	}
	for i, r := range suf {
		permutate(pref+string(r), suf[:i]+suf[i+1:], p)
	}
}
