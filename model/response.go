package model

// Response representa o formato da resposta da API
type Response struct {
	Source  string `json:"source"`
	Address string `json:"address"`
	Error   string `json:"error,omitempty"`
}
