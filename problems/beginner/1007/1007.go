package main

import "fmt"

func main() {
	var A, B, C, D, dif int
	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)
	fmt.Scan(&D)

	dif = (A*B - C*D)

	fmt.Printf("DIFERENCA = %d\n", dif)
}
