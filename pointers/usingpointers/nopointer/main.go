package main

import "fmt"

/* func zero(z int) {
	z = 0
}

func main() {
	x := 5
	zero(x)
	fmt.Println(x) // value 5 is passed to func zero, so the output is 5, x = 0 has no effect
}
*/

func zero(x int) {
	fmt.Println(&x) // address in func zero
	x = 0
}

func main() {
	y := 5
	fmt.Println(&y) // address in main
	zero(y)
	fmt.Println(y) // y is still 5
}
