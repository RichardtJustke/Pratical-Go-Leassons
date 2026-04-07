package main

import (
	"fmt"
)

func main() {
	n := 2548
	n2 := 0x9F4
	n3 := 02454
	fmt.Printf("Hexadecimal: %x\n", n) // o %x e o formatador para nmero hexadecimais
	fmt.Printf("Decimal: %d\n", n2)    // o \n serve para que tenha a quebra de linha
	fmt.Printf("Octal: %o\n", n3)      // formatador %o para numeros octais
	fmt.Printf("Colocando prefixo: %O\n", n3)
	// "%o"permitir imprimir o número no octal
	//"%O"permite imprimir o número em octal com um "0o"prefixo

	/*
		1. Formatação Automática de Código (gofmt)
		Go impõe um estilo único e padronizado, eliminando debates sobre estilo.

		    gofmt: Ferramenta padrão que formata automaticamente o código Go.
		    Indentação: Usa tabulações (tabs) em vez de espaços por padrão.
		    Regras: Chaves na mesma linha, fim de linha sem ponto e vírgula, organização automática de imports.

		2. Formatação de Strings e Dados (Pacote fmt)
		O pacote fmt é usado para formatar saídas, logs e strings usando verbos (placeholders).
		Verbos Principais (fmt.Printf, fmt.Sprintf):

		    %v: O valor no formato padrão (vantagem: funciona para quase tudo).
		    %+v: Ao imprimir structs, adiciona os nomes dos campos.
		    %#v: Representação sintática Go do valor (útil para depuração).
		    %T: O tipo do valor.
		    %%: Um sinal de porcentagem literal.

		Verbos Numéricos e de Texto:

		    %t: Booleano (true ou false).
		    %d: Inteiro em base 10.
		    %b: Base 2 (binário).
		    %x / %X: Base 16 (hexadecimal).
		    %f / %F: Ponto flutuante sem expoente (ex: 123.456).
		    %s: String bruta ou slice de bytes.
		    %q: String com aspas duplas seguras.

		Verbos de Ponteiro:

		    %p: Endereço de memória em hexadecimal.

		3. Formatação de Tempo e Datas
		Go utiliza um método único baseado em um layout de referência específico: 01/02 03:04:05PM '07 -0700.

		    Exemplo: data.Format("2006-01-02") para formatar no estilo AAAA-MM-DD.

		4. Principais Funções de Formatação (fmt)

		    fmt.Print: Imprime sem formatação especial.
		    fmt.Println: Imprime com nova linha no final.
		    fmt.Printf: Formata de acordo com o verbo e imprime.
		    fmt.Sprintf: Retorna a string formatada sem imprimi-la.
	*/
}
