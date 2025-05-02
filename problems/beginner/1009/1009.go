package main

import "fmt"

func main() {
	var nome string
	var salarioFixo, totalVendas, salarioTotal float64

	fmt.Scan(&nome)
	fmt.Scan(&salarioFixo)
	fmt.Scan(&totalVendas)

	salarioTotal = (salarioFixo + totalVendas*0.15)

	fmt.Printf("TOTAL = R$ %.2f\n", salarioTotal)
}
