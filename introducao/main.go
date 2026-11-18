// o pacote main não pode ser importado. Ele é a base, utilizado apenas para criar executáveis!
package main

import (
	"fmt"
	"introducao/pacote"
)

func main() {
	fmt.Println("Hello, world!")

	fmt.Println(pacote.Bar)
	fmt.Println(pacote.Foo)
	pacote.PrintMinha()
}
