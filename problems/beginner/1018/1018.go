package main

import "fmt"

func main() {
	var valor, troco, cem, cinquenta, vinte, dez, cinco, dois, um int
	fmt.Scan(&valor)
	troco = valor
	cem = troco / 100
	troco %= 100
	cinquenta = troco / 50
	troco %= 50
	vinte = troco / 20
	troco %= 20
	dez = troco / 10
	troco %= 10
	cinco = troco / 5
	troco %= 5
	dois = troco / 2
	troco %= 2
	um = troco
	fmt.Println(valor)
	fmt.Println(cem, "nota(s) de R$ 100,00")
	fmt.Println(cinquenta, "nota(s) de R$ 50,00")
	fmt.Println(vinte, "nota(s) de R$ 20,00")
	fmt.Println(dez, "nota(s) de R$ 10,00")
	fmt.Println(cinco, "nota(s) de R$ 5,00")
	fmt.Println(dois, "nota(s) de R$ 2,00")
	fmt.Println(um, "nota(s) de R$ 1,00")
}
