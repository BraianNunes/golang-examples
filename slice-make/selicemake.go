package main

import "fmt"

func main() {
	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)

	s = make([]int, 10, 20)

	// len(s) = tamanho
	// cap(s) = capacidade do array
	// disponibiliza 10 elementos mas o array possui capacidade para 100 elementos
	fmt.Println(s, len(s), cap(s))

	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(s, len(s), cap(s))

	// Se preencher um slice acima da capacidade dele o Golang duplica a capacidade
	s = append(s, 10)
	fmt.Println(s, len(s), cap(s))


}
