package main

import "fmt"

func main() {
	v := 42.0 // type inference works in one direction only. can change only in this line
	fmt.Printf("v is of type %T\n", v)
}
