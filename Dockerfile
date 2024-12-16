# Etapa 1: Build
FROM golang:1.23.4 AS builder

WORKDIR /app

# Copiar os arquivos do módulo
COPY go.mod go.sum ./

# Baixar dependências do repositório Git
RUN go mod tidy

# Copiar o restante do código
COPY . .

# Compilar o binário
RUN go build -o app .

# Etapa 2: Runtime
FROM debian:bullseye-slim

WORKDIR /app

# Copiar o binário gerado no estágio de build
COPY --from=builder /app/app .

CMD ["./app"]