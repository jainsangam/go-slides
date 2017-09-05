package main
import (
	"fmt"
	"strconv"
)
func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Print(strconv.Itoa(i)+" ")
	}
}
/* output
counting
9 8 7 6 5 4 3 2 1 0
*/
