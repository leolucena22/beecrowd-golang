package main

import "fmt"

func main() {
	var distancia, velocidade, tempo int
	var gasto float64
	fmt.Scan(&velocidade)
	fmt.Scan(&tempo)

	distancia = velocidade * tempo

	gasto = float64(distancia) / 12.0

	fmt.Printf("%.3f\n", gasto)
}
