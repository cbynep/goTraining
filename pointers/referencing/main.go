package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a) // 43
	fmt.Println(&a)

	var b *int = &a // b is a pointer to an int, it points where a is stored, *int is a type, b references a's memory address

	fmt.Println(b)
	fmt.Println(*b) // 43, b points to a memory address, so * dereferences that to show value and in this case * is an operator

	*b = 42        // b changes the value at this address to 42
	fmt.Println(a) // 42

}
