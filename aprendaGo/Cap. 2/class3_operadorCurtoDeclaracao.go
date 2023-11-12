package main

/*Operador Curto de Declaração := 'Marmota' (Gopher) -> mascote do GO -> só pode ser usada dentro de codeBlock*/
/*Tipagem automática */

import "fmt"

/*
Escopos - pode ser a nível de packages (global, como alguns falam) ou então de função (local)
Vale ressaltar que o as variáveis só podem ser acessadas após seu escopo, ou seja, trocar o a linha 25 para ficar
antes da linha 24 (declaração da variável) ocasionária erro pelo fato da variável não abrangir aquele escopo
*/
var y = "Good morning"

/*key-words -> palavras reservadas. Só podem ser usadas para seu devido funcionamento, não podendo ser utilizadas
em outros contextos */

func main() {
	x := 8 + 3

	/*O 'format print' (Printf) ele vem de origem da linguagem C, logo o '%v' mostra o valor da variável e a
	'%T' mostra seu tipo */
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n\n", y, y)

	/*Quando utilizamos o ':=' significa que estamos declarando uma variável, quando queremos apenas atribuir um
	novo valor devemos apenas usar '=' para receber. A marmota (:=) precisa ao menos declarar uma nova variável
	a esquerda para poder ser compilada, e ela ja entende qual deve ser atribuida e qual deve ser declarada */
	/*Operador de comparação -> '==' */
	x, z := 16.0, 5 == 5
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
}
