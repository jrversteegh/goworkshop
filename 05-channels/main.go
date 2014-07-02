package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
  c := make(chan string)
	go talk("rob", c)
	go talk("christian", c)

	time.Sleep(10 * time.Second)
}

func talk(name string, channel chan string) {
	for {
    s := fmt.Sprintf("my name is %v\n", name)
    channel <- s
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	}
}
