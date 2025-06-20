package main

import (
	"fmt"
)

func main() {
	var a, b, c, maior int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	if a > b && a > c {
		maior = a
	} else if b > a && b > c {
		maior = b
	} else if c > a && c > b {
		maior = c
	}

	fmt.Printf("%d eh o maior\n", maior)
}
