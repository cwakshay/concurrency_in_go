package main

import (
	"fmt"
	"time"
)

func main() {
	for x := 0; x < 1000; x++ {
		call()
		time.Sleep(2 * time.Millisecond)
		fmt.Println()
	}
}
func call() {
	var x, y int
	go func() {
		x = 1
		fmt.Print("y: ", y, " ")
	}()

	go func() {
		y = 1
		fmt.Print("x: ", x, " ")
	}()

}
