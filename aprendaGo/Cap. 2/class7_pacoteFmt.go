package main

import "fmt"

/*
String: interpreted string literals vs. raw string literals
Literal é uma notação para representar um valor fixo no código fonte
*/

/*
Format printing: documentação -> FMT
Todas possuem o seu comum (print), depois adicionar uma linha ao final (println) e depois para exibir
o formato (printf)
Print comum -> serve para exibir o dado para o usuário
Sprint serve para apenas armazenar a string, ele não será enviado para a tela
Fprint serve para files basicamente
*/

func main() {
	x := "Hello, good morning\nHow are you?\nI hope everything is fine!"
	fmt.Println(x)

	y := "oi"
	z := "eita"
	w := fmt.Sprint(y, " ", z)

	fmt.Println(w)
}