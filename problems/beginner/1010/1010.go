package main

import "fmt"

func main() {
	var codPeca1, codPeca2, qtdPecas1, qtdPecas2 int
	var valorPeca1, valorPeca2, total float64

	fmt.Scanf("%d %d %f", &codPeca1, &qtdPecas1, &valorPeca1)

	fmt.Scanf("%d %d %f", &codPeca2, &qtdPecas2, &valorPeca2)

	total = float64(qtdPecas1)*valorPeca1 + float64(qtdPecas2)*valorPeca2

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
