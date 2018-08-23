package main

import "fmt"

//List 200 numbers in dec, bin, hex
/* func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#x \n", i, i, i)
	}
}
*/

//Plus %q for UTF-8
func main() {
	for i := 0; i < 200; i++ {
		fmt.Printf("%d \t %b \t %#x \t %q \n", i, i, i, i)
	}
}
