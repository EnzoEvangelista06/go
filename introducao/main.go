// Nada pode vir antes da declaração do pacote.
package main

import (
	// renomeação
	meuPacote "encoding/json"
	// Utilizando o ".", é como se o pacote fosse "incorporado" em nosso sistema.
	// dot import --> não recomendado
	. "fmt"

	// Isso existe. Ele não entra no escopo do nosso pacote atual.
	_ "fmt"
)

func main() {
	// fmt.Println("Olá, mundo!")
	Println("Hello, world!")
	meuPacote.Marshal()
}
