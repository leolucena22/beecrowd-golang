package main

import "fmt"

func main() {
	var A, B, media float64

	fmt.Scan(&A)
	fmt.Scan(&B)

	media = (A*3.5 + B*7.5) / 11

	fmt.Printf("MEDIA = %.5f\n", media)
}
