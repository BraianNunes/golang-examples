package main

import "fmt"

func f1() {
	fmt.Println("teste")
}

func f2(p1 string, p2 string) {
	fmt.Printf("F2: %s %s\n", p1, p2)
}

func f3() string {
	return "F3"
}

func f4(p1, p2 string) string {
	return fmt.Sprintf("F4: %s %s", p1, p2)
}

func f5() (string, string) {
	return "Olá", "Mundo"
}

func f6() (int, int, int, int) {
	return 1,2,3,4
}

func main() {
	f1()

	f2("Olá", "Mundo")
	f3_ := f3()
	fmt.Println(f3_)

	r3, r4 := f3(), f4("Param1", "Param2")
	fmt.Println(r3, r4)

	retorno1, retorno2 := f5()
	fmt.Println(retorno1, retorno2)

	fmt.Println(f6())
}