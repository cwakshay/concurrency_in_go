package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(1 * time.Second)
	fmt.Println("Count down begins")

	for x := 10; x > 0; x-- {
		fmt.Println(x)
		<-tick
	}

	fmt.Println("Launch...")
}
