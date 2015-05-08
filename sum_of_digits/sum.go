package main

import "fmt"

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
		//'scanner.Text()' represents the test case, do something with it
		fmt.Println(sum(scanner.Text()))
	}
}

func sum(in string) int {
	result := 0
	for i := 0; i < len(in); i++ {
		result += int(in[i]) - '0'
	}
	return result
}
