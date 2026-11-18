package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello, world!")

	// chamada de função
	digaOi()
	fmt.Println(somar(5, 2))

	a, b := swap(10, 20)
	fmt.Println(a, b)

	// divisão
	res, rem := dividir(10, 3)
	fmt.Println(res, rem)

	// higher order
	x := higherOrder(10)(5)
	fmt.Println(x)

	// função variática (diversos parâmetros)
	fmt.Println(variaticaSomar(10, 10, 10))
}

// função + nome
func digaOi() {
	fmt.Println("Oi")
}

// parâmetros e retorno tipados
func somar(a int, b int) int {
	return a + b
}

// uma função pode retornar múltiplos valores
func swap(a, b int) (int, int) {
	return b, a
}

// podemos nomear os dois retornos
func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b

	return res, rem
}

// função Higher Order
func higherOrder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// função variática
func variaticaSomar(nums ...int) int {
	var out int
	for _, n := range nums {
		out += n
	}

	return out
}
