package main

import (
	"fmt"
)

func main() {

	var A, B, X int

	fmt.Scanln(&A)
	fmt.Scanln(&B)

	X = A + B

	fmt.Printf("X = %d\n", X)

}
