package main

import (
	"fmt"

	"time"
)

func main() {
	c, b := make(chan int), make(chan rune)
	go func() {
		for i := 1; i < 26; i++ {
			c <- i
			// fmt.Print(i)

		}
	}()
	go func() {
		for j := 'A'; j <= 'z'; j++ {
			b <- j
		}
	}()
	go func() {
		for {
			fmt.Print(<-c)
			fmt.Print(string(<-b))

		}
	}() 
	time.Sleep(time.Second)

}
