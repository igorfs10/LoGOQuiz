package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var logos []bool
var quantidadeLogo int

func main() {
	reader := bufio.NewReader(os.Stdin)
	quantidadeLogo = PegarQuantidadeLogos()
	for i := 0; i < quantidadeLogo; i++ {
		logos = append(logos, false)
	}
	fmt.Println("My favorite number is", rand.Intn(10), "na hora", time.Now())
	var logo Logo = PegarLogo(1)
	//ImprimirLogo(logo.Caminho)
	// 0 = 0, 1 = 2, 2 = 5, 3 = 7
	var nome string = NomeParaLinhas(logo.Nome)

	ImprimirLogo(logo.Caminho)

	fmt.Println(logos)
	fmt.Println(logo.Nome)

	fmt.Println(nome)

	nome, _ = UsarDica(nome, logo.Nome)
	fmt.Println(nome)

	fmt.Println(NormalizaStr("Olá mundo-novo"))

	var storageString string
	storageString, _ = reader.ReadString('\n')
	fmt.Println(storageString)

	printMenu()

	var language LanguageItens
	CarregaIdioma(&language)
}

/*------------- Tratar para retornar um erro se não possuir carácter para substituir -----------*/

// UsarDica : Muda uma linha do nome para a letra do nome
func UsarDica(nomeLinha string, nome string) (string, bool) {
	// Verifica se a string tem "_" para substituir
	if strings.Contains(nomeLinha, "_") {
		var numChar int
		var numLinha int
		for {
			rand.Seed(time.Now().UnixNano())
			// Intn retorna um inteiro entre 0 e n -1
			numChar = rand.Intn(len(nome))
			numLinha = numChar * 2
			if string(nomeLinha[numLinha]) == "_" {
				break
			}
		}
		return ReplaceAtIndex(nomeLinha, rune(nome[numChar]), numLinha), true
	}
	// Tratar para retornar um erro se não estiver carácter para substituir
	return nomeLinha, false
}

/*------------- Tratar para retornar um erro se não possuir carácter para substituir -----------*/

func printMenu() {
	fmt.Println("——————————————————————————————————————————————————————————————————————————————————————————————————")
	fmt.Println(" 1 - $TXT_JOGAR      2 - $TXT_SAIR")
	fmt.Println("——————————————————————————————————————————————————————————————————————————————————————————————————")
	// %v imprime no formato do item passado
	fmt.Printf("%v ---\n", "test")
}
