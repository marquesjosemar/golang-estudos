package main

import "fmt"

func main() {
	// A nossa "casa" com o valor 42 dentro.
	numero := 42

	// 'p' agora guarda o ENDEREÇO da variável 'numero'.
	// 'p' é um ponteiro para um inteiro. O tipo dele é *int.
	p := &numero

	/*
		p não contém 42, ele contém o endereço 0xc000018088, que é onde o 42 está guardado na memória do computador. */

	fmt.Println("Valor original de 'numero':", numero)
	fmt.Println("Endereço de 'numero' guardado em 'p':", p)
}
