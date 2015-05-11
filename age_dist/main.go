package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		fmt.Println(dist(scanner.Text()))
	}
}

func dist(s string) string {
	i, _ := strconv.Atoi(s)
	if i < 0 || i > 100 {
		return "This program is for humans"
	} else if i <= 2 {
		return "Still in Mama's arms"
	} else if i <= 4 {
		return "Preschool Maniac"
	} else if i <= 11 {
		return "Elementary school"
	} else if i <= 14 {
		return "Middle school"
	} else if i <= 18 {
		return "High school"
	} else if i <= 22 {
		return "College"
	} else if i <= 65 {
		return "Working for the man"
	} else if i <= 100 {
		return "The Golden Years"
	}
	return ""
}
