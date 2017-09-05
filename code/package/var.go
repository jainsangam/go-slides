package main

import "fmt"

var h int = 4
func main() {

	var i int 		// declaration
	i = 3
	var j int = 12	// declaration with initialization
	k := 3			// short initialization ( k is of type int )

	c, python, java := true, false, "no!"	// multiple initialization

	fmt.Println(h, i, j, k, c, python, java)	//output: 4 3 12 3 true false no!
}
