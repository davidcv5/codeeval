package main

import (
	"fmt"
	"log"
)

import "os"

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	i, err := file.Stat()
	if err != nil {
		fmt.Println(i.Size())
	}
}
