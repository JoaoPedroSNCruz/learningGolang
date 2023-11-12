package main

import "fmt"

type hotdog int

var b hotdog = 101

func main() {
	x := 10
	fmt.Printf("%v\n", x)

	/*Conversão -> alteração de tipos e usabilizade de outras variáveis sem ser o mesmo tipo 
	Basicamente para fazer deve-se colocar o 'tipo(var)' para obter o valor convertido */
	x = int(b)
	fmt.Printf("%v\n", x)
}