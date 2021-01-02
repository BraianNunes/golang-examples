package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64 {
		"José João": 11265.54,
		"Gabrielle Santos": 15000.00,
		"Pedro Junior": 1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.00
	fmt.Println(funcsESalarios)

	// Não há problemas em excluir chaves que não existem
	delete(funcsESalarios, "Inexistente")
}
