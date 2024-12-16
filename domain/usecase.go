package domain

import (
	"context"
	"errors"
	"time"

	"github.com/robertocorreajr/fullcycle-multithreading-challenge/infrastructure"
	"github.com/robertocorreajr/fullcycle-multithreading-challenge/model"
)

// GetAddress é o caso de uso para buscar o endereço pelo CEP
func GetAddress(cep string) (model.Response, error) {
	// URLs das APIs
	api1 := "https://brasilapi.com.br/api/cep/v1/" + cep
	api2 := "http://viacep.com.br/ws/" + cep + "/json/"

	ch := make(chan model.Response, 2)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// Requisições concorrentes
	go infrastructure.FetchAPI(ctx, api1, "BrasilAPI", ch)
	go infrastructure.FetchAPI(ctx, api2, "ViaCEP", ch)

	// Captura a resposta mais rápida
	select {
	case res := <-ch:
		if res.Error != "" {
			return model.Response{}, errors.New(res.Error)
		}
		return res, nil
	case <-ctx.Done():
		return model.Response{}, errors.New("timeout: nenhuma resposta no tempo limite")
	}
}
