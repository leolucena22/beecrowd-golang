package main

import "fmt"

func main() {
	var raio, area float64

	fmt.Scanln(&raio)

	area = 3.14159 * (raio * raio)

	fmt.Printf("A=%.4f\n", area)
}
