package main

import (
	"fmt"
)

//Regular print line
/* func main() {
	fmt.Println(42)
} */

//Decimal, binary Printf
/* func main() {
	fmt.Printf("%d - %b", 42, 42)
} */

//Decimal, binary, hexadecimal
/*  func main() {
	fmt.Printf("%d - %b - %x", 42, 42, 42)
 } */

//With format verbs %#x for hex, %#o for octa and escape characters, \n - new line, \t - tab
func main() {
	fmt.Printf("%d \t %b \t %#o \t %#x \n", 42, 42, 42, 42)
}
