package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(roman(scanner.Text()))
	}
}

func roman(s string) string {
	n := len(s)
  for 
}

func digit(s string, p int) string{
  x,_:= strconv.Atoi(s)
  if x != 4{

  }
}

func letter(p int) string{
  switch p{
  case 1:
    return "I"
    case 2:
    return "X"
    case 3:
    return "C"
    case 4:
    return "M"
  }
}