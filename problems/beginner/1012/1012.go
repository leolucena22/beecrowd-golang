package main

import "fmt"

func main() {
	var A, B, C, trianguloRet, circulo, trapezio, quadrado, retangulo float64

	fmt.Scanf("%f %f %f", &A, &B, &C)

	trianguloRet = (A * C) / 2
	circulo = 3.14159 * (C * C)
	trapezio = (A + B) * C / 2
	quadrado = B * B
	retangulo = A * B

	fmt.Printf("TRIANGULO: %.3f\nCIRCULO: %.3f\nTRAPEZIO: %.3f\nQUADRADO: %.3f\nRETANGULO: %.3f\n", trianguloRet, circulo, trapezio, quadrado, retangulo)
}
