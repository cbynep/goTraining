package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {

	a := 43

	fmt.Println("a -", a)
	fmt.Println("a's memory address -", &a)
	fmt.Printf("decimal - %d\n", &a)

	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters) //takes user input
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")

}
