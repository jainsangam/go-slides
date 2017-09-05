package main

import "fmt"

var power = []int{1, 2, 4, 8 }

func main() {
	for i, v := range power {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
/** output
2**0 = 1
2**1 = 2
2**2 = 4
2**3 = 8
 */
