package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/robertocorreajr/fullcycle-multithreading-challenge/domain"
)

func main() {
	// Captura o CEP como argumento de linha de comando
	cep := flag.String("cep", "", "CEP a ser consultado (exemplo: 01153000)")
	flag.Parse()

	if *cep == "" {
		fmt.Println("Erro: O parâmetro 'cep' é obrigatório. Use: ./app -cep=01153000")
		os.Exit(1)
	}

	// Executa o caso de uso
	result, err := domain.GetAddress(*cep)
	if err != nil {
		fmt.Printf("Erro: %s\n", err.Error())
		return
	}

	fmt.Printf("Resultado: %+v\n", result)
}
