package main

import (
	"fmt"
	"math/rand"
	"time"
)

var logos []bool
var quantidadeLogo int

func main() {
	quantidadeLogo = PegarQuantidadeLogos()
	for i := 0; i < quantidadeLogo; i++ {
		logos = append(logos, false)
	}
	fmt.Println("My favorite number is", rand.Intn(10), "na hora", time.Now())
	var logo Logo = PegarLogo(0)
	//ImprimirLogo(logo.Caminho)
	// 0 = 0, 1 = 2, 2 = 5, 3 = 7
	var nome string = NomeParaLinhas(logo.Nome)

	ImprimirLogo(logo.Caminho)

	fmt.Println(logos)
	fmt.Println(logo.Nome)

	fmt.Println(nome)

	nome = UsarDica(nome, logo.Nome)
	fmt.Println(nome)

	nome = UsarDica(nome, logo.Nome)
	fmt.Println(nome)

	nome = UsarDica(nome, logo.Nome)
	fmt.Println(nome)

}
