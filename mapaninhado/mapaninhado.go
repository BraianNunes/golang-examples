package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 15000.00,
			"Guga Pereira": 10000.00,
		},
		"J": {
			"João Silva": 9000.00,
			"José Aguiar": 11000.00,
		},
		"P": {
			"Pedro Miguel": 16000.00,
		},
	}
	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Printf("%s -- %s", letra, funcs)
	}

}
