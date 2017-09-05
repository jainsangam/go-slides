package main

import "fmt"

func forDefault() (sum int) {
	for i := 0; i < 10; i++ {
		sum += i
	}
	return sum
}

func forOptional() (sum int) {
	i := 0
	for ; i < 10 ; { 	// optional init and post statement
// similar to for i < 10  {	// for as while
		sum += i;
		i += 1
	}
	return sum
}
func main() {
	fmt.Println(forDefault())	//output : 45
}