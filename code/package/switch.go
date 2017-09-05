package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {	// empty condition can be replaced with switch _some_expression_ {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
