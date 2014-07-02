package main

import "fmt"
import "time"
import "math/rand"

func main() {
  rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("My favorite number is", rand.Intn(11))
}
