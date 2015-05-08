package main

import (
	"errors"
	"fmt"
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
		s := &stack{}
		list := strings.Fields(scanner.Text())
		for _, v := range list {
			s.push(v)
		}
		line := make([]string, 0)
		for {
			if x, err := s.pop(); err == nil {
				line = append(line, x)
			} else {
				break
			}
		}

		fmt.Println(strings.Join(line, " "))
	}
}

type stack struct {
	values []string
}

func (s *stack) push(i string) {
	s.values = append(s.values, i)
}

func (s *stack) pop() (string, error) {
	if len(s.values) == 0 {
		return "", errors.New("Empty")
	}

	r := s.values[len(s.values)-1]
	if len(s.values) == 1 {
		s.values = s.values[1:]
	} else {
		s.values = s.values[:len(s.values)-2]
	}
	return r, nil
}
