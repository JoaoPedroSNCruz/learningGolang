package main

/*Sempre que houver na documentação '...' antes de uma 'interface()' significa que aquela função pode receber
N parametros e de N tipos */

import "fmt"

func main() {
	/*A função 'Pintln' que está localizada pela package 'format' (fumit) ela além de funcionar como um output
	ela também pode retornar o 'número de bytes' que o que ela está imprimindo possui e pode demonstrar a
	quantidade de 'erro's ao imprimir */
	numeroBytes, erros := fmt.Println("Hello, wordl", "Salve guys", 19)
	fmt.Println(numeroBytes, erros)

	/*A linguagem Go ela entende que todo o código deve ser performático e eficiente, por isso caso você tenha
	uma variável sem uso em seu código ela não permite que o código seja compilado, logo para evitar situações
	que podem ocorrer caso uma função te retorne por exemplo 2 valores porém você deseja apenas 1 podemos usar o
	'_' (underscore) que é conhecido também como 'blankIdentifier' */
	blankIdentifier, _ := fmt.Println("Hello, wordl", "Salve guys", 19)
	fmt.Println(blankIdentifier)

	var x int = 8
	/*String sempre será dentro de "" e não '' */
	var y string = "strings"
	/*bool -> booleano -> True/False */
	var z bool = false

	fmt.Println(x, y, z)
}
