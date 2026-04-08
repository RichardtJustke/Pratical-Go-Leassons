package main

import "fmt"

func main() {
	var nome string = "Justke" // variavel tipada
	idade := 20

	const PI = 3.14 // constante imutavel

	fmt.Println(nome, idade, PI)
}

/*
 	vairavel sempre pode mudar e tem tipo fixo (usar quando o valor muda inputs, contadores,etc...)
	constante não pode mudar e nem sempre tem tipo (usar quando tem valores fixos PI, status, config)

*/
