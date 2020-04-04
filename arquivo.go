package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// JSONParaVar : Carrega um arquivo JSON e coloca numa variável. Ele usa o ponteiro da variável e o caminho do arquivo JSON.
func JSONParaVar(variavel interface{}, caminho string) {
	jsonFile, _ := os.Open(caminho)
	defer jsonFile.Close()
	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &variavel)
}

// PegarLogo : Pega um logo no arquivo JSON para usar no programa
func PegarLogo(ID int) Logo {
	var logos []Logo
	JSONParaVar(&logos, "logos.json")
	return logos[ID]
}

// ImprimirLogo : Mostra o logo na tela usando o nome do arquivo como parâmetro
func ImprimirLogo(arquivoLogo string) {
	file, _ := os.Open("logos/" + arquivoLogo)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

// PegarQuantidadeLogos : Pega a quantidade de logos que existem
func PegarQuantidadeLogos() int {
	var logos []Logo
	JSONParaVar(&logos, "logos.json")
	return len(logos)
}
