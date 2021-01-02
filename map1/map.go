package main

import "fmt"

func main() {
	// Mapas devem ser incializados

	aprovados := make(map[int]string)
	aprovados[12345678] = "Maria"
	aprovados[54216898] = "Pedro"
	aprovados[12254882] = "Ana"

	fmt.Println(aprovados)


	// for
	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	// Print by key
	fmt.Println(aprovados[12345678])

	// Remove by key
	delete(aprovados, 12345678)
	fmt.Println(aprovados[12345678])
}
