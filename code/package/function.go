package main

import "fmt"
// named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
// multiple return value
func swap(x, y string) (string, string) {
	return y, x
}
// simple function
func add (x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(1,3));	// output: 4
	fmt.Println(swap("world","hello"))	//output: hello world
	fmt.Println(split(17))		//output: 7 10
}

