package main

import "fmt"

func main() {
	raw := `spring rain:
browsing under an umbrella
at the picture-book store`
	fmt.Println(raw) // quando usamos crase em um citação transoformamos ela em uma string literal

	interpreted := "i love spring"
	fmt.Println(interpreted)
}

/*
Existem dois “tipos” de literais de strings:

    literais de corda cru. São definidas entre citações traseiras.

        Personagens proibidos são
            back citações

        Personagens descartados são
            Devoluções de transporte (\r)

    interpretedinterpretados string literais. São definidas entre as quotas duplas.

        Personagens proibidos são

            novas linhas

            citações duplas não escapadas
*/
