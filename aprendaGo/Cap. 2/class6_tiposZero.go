package main

import "fmt"

/*
Declarar -> 'criar' a variável
Inicializar -> Primeiro valor da variável
Atribuir -> Novo valor da variável
Marmota (:=) faz tudo 
*/

/*Cada tipo possui um zero especifico
Para pointes, functions, interfaces, slices, channels, maps: nul */

/*
Sempre que possível use (:=) 'marmota'
Use 'var' para package-level scope
*/

var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}
