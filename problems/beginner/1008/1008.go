package main

import "fmt"

func main() {
	var numFuncionario, numHorasTrabalhadas int
	var valorHora, salarioTotal float64

	fmt.Scan(&numFuncionario)
	fmt.Scan(&numHorasTrabalhadas)
	fmt.Scan(&valorHora)

	salarioTotal = float64(numHorasTrabalhadas) * valorHora

	fmt.Printf("NUMBER = %d\nSALARY = U$ %.2f\n", numFuncionario, salarioTotal)
}
