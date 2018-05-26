package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10000; i++ {
		go func() {
			fmt.Print(i)
			fmt.Println("\n")
		}()
		time.Sleep(1 * time.Millisecond)
	}

	time.Sleep(1 * time.Second)
}
