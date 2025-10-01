package main

import "fmt"

func main() {
	soma, subtracao, multiplicacao, divisao := Operacao(1, 2)

	fmt.Println(soma, subtracao, multiplicacao, divisao)
}

func Operacao(numero1 int, numero2 int) (int, int, int, int) {
	soma := numero1 + numero2
	subtracao := numero1 - numero2
	multiplicacao := numero1 * numero2
	divisao := numero1 / numero2

	return soma, subtracao, multiplicacao, divisao
}
