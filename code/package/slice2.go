package main

import "fmt"
func main() {
	s := make([]int, 5);printSliceElement(s)
	// append works on nil slices.
	s = append(s, 0);printSliceElement(s)
	// The slice grows as needed.
	s = append(s, 1);printSliceElement(s)
	// We can add more than one element at a time.
	s = append(s, 2, 3, 4);printSliceElement(s)
}
func printSliceElement(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
/* output
len=5 cap=5 [0 0 0 0 0]
len=6 cap=10 [0 0 0 0 0 0]
len=7 cap=10 [0 0 0 0 0 0 1]
len=10 cap=10 [0 0 0 0 0 0 1 2 3 4]
 */