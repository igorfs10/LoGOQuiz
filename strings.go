package main

import (
	"math/rand"
	"reflect"
	"time"
	"unsafe"
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

// UsarDica : Muda uma linha do nome para a letra
func UsarDica(nomeLinha string, nome string) string {
	rand.Seed(time.Now().UnixNano())
	// Intn retorna um inteiro entre 0 e n -1
	var numChar int = rand.Intn(len(nome))
	var numLinha int = numChar * 2
	for string(nomeLinha[numLinha]) != "_" {
		rand.Seed(time.Now().UnixNano())
		numChar = rand.Intn(len(nome))
		numLinha = numChar * 2
	}
	return replaceAtIndex(nomeLinha, rune(nome[numChar]), numLinha)
}

func strToBytes(str string) []byte {
	stringHeader := (*reflect.StringHeader)(unsafe.Pointer(&str))
	bytesHeader := &reflect.SliceHeader{
		Data: stringHeader.Data,
		Len:  stringHeader.Len,
		Cap:  stringHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(bytesHeader))
}

func bytesToStr(bytes []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&bytes))
	stringHeader := &reflect.StringHeader{
		Data: bytesHeader.Data,
		Len:  bytesHeader.Len,
	}
	return *(*string)(unsafe.Pointer(stringHeader))
}

func replaceAtIndex(str string, replacement rune, index int) string {
	bytes := strToBytes(str)
	bytes[index] = byte(replacement)
	return bytesToStr(bytes)
}
