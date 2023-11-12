package main

import "fmt"

type hotdog int

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v %T", b, b)

	/*
		x = b
		Não funciona porque ja que são tipos diferentes não existe nenhum comparação possível
	*/
}
