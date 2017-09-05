package main

import (
	"fmt"
	"math"
)
func pow(x, n, lim float64) float64 {
	/*
		v := math.Pow(x,n);
	 	if v < lim { 	*/
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),	//output:  27 >= 20
		pow(3, 3, 20),	//output: 9 20
	)
}
