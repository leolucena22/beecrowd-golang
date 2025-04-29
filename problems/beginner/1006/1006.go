package main

import "fmt"

func main() {
	var A, B, C, media float64

	fmt.Scan(&A)
	fmt.Scan(&B)
	fmt.Scan(&C)

	media = (A*2 + B*3 + C*5) / 10

	fmt.Printf("MEDIA = %.1f\n", media)
}
