package main

import (
	"fmt"
	operacoes_array "github.com/ericpatricksantos/operacoes_array"
)

func main() {
	lista1 := []string{
		"1", "1", "2",
	}

	list2 := []string{
		"3", "4", "4",
	}

	listaR, mapa, tam := operacoes_array.UneDuasListasSemElementosRepetidos(lista1, list2)

	fmt.Println("Lista String: ", listaR)

	fmt.Println("----------------------")

	fmt.Println("Mapa: ", mapa)

	fmt.Println("----------------------")

	fmt.Println("tamanho da Lista e do Mapa:", tam)
}
