package main

import (
	"fmt"
)

func main() {
	var x int
	var y, consumo float64

	fmt.Scan(&x)
	fmt.Scan(&y)

	consumo = float64(x) / y
	fmt.Printf("%.3f km/l\n", consumo)

}
