package infrastructure

import (
	"context"
	"io"
	"net/http"

	"github.com/robertocorreajr/fullcycle-multithreading-challenge/model"
)

// FetchAPI faz a requisição para uma API específica
func FetchAPI(ctx context.Context, url, source string, ch chan<- model.Response) {
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	client := http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		ch <- model.Response{Source: source, Error: "erro ao conectar na API"}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ch <- model.Response{Source: source, Error: "erro ao ler a resposta da API"}
		return
	}

	ch <- model.Response{Source: source, Address: string(body)}
}
