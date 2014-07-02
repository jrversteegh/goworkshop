package main

import "fmt"

func even(value int) (bool) {
  return (value % 2) == 0
}

func main() {
	n := 13

	if even(n) {
		fmt.Printf("%v is even!\n", n)
	} else {
		fmt.Printf("%v is odd!\n", n)
	}
}
