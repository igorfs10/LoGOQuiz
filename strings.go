package main

import (
	"math/rand"
	"reflect"
	"strings"
	"time"
	"unicode"
	"unsafe"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// NomeParaLinhas : converte o nome para o nome em linhas
func NomeParaLinhas(nome string) string {
	var tamanhoTexto int = len(nome)
	var nomeLinha string
	for i := 0; i < tamanhoTexto; i++ {
		if string(nome[i]) == "-" {
			nomeLinha += "-"
		} else {
			nomeLinha += "_"
		}

		if i < tamanhoTexto-1 {
			nomeLinha += " "
		}
	}
	return nomeLinha
}

/*------------- Tratar para retornar um erro se não possuir carácter para substituir -----------*/

// UsarDica : Muda uma linha do nome para a letra do nome
func UsarDica(nomeLinha string, nome string) string {
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
		return replaceAtIndex(nomeLinha, rune(nome[numChar]), numLinha)
	}
	// Tratar para retornar um erro se não estiver carácter para substituir
	return nomeLinha
}

/*------------- Tratar para retornar um erro se não possuir carácter para substituir -----------*/

// Converte string em bytes
func strToBytes(str string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bytesHeader := &reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(bytesHeader))
}

// Converte bytes em string
func bytesToStr(bytes []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	stringHeader := &reflect.StringHeader{
		Data: bytesHeader.Data,
		Len:  bytesHeader.Len,
	}
	return *(*string)(unsafe.Pointer(stringHeader))
}

// Muda um carácter passando o indíce do vetor usando as funções de conversão de string para uma performance maior
func replaceAtIndex(str string, replacement rune, index int) string {
	bytes := strToBytes(str)
	bytes[index] = byte(replacement)
	return bytesToStr(bytes)
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// Função para remover acento das letras
func removeAccent(str string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	result, _, _ := transform.String(t, str)
	return result
}

// NormalizaStr : retorna uma string sem espaços e caracteres especiais
func NormalizaStr(str string) string {
	var minuscula string = strings.ToLower(str)
	minuscula = removeAccent(minuscula)

	// O Replacer troca todos os caracteres dentro da string, sempre trocando o carácter do parâmetro trocando pelo caracter seguinte
	var replacer = strings.NewReplacer(" ", "", "-", "")

	return replacer.Replace(minuscula)
}
