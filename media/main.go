package main

import "fmt"

func main() {

	lista := make([]int, 1)
	lista[0] = 1
	lista = append(lista, 2)
	lista = append(lista, 3)

	somaTotal := 0
	for i := 0; i < len(lista); i++ {
		somaTotal += lista[i]
	}

	fmt.Println("Média: ", somaTotal/len(lista))
}
