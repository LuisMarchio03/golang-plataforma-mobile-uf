# Estágio de build
FROM golang:1.24.0-alpine AS builder

WORKDIR /app

# Adiciona dependências necessárias para build
RUN apk add --no-cache gcc musl-dev

# Copia os arquivos go.mod e go.sum
COPY go.mod go.sum ./

# Download das dependências
RUN go mod download

# Copia o código fonte
COPY . .

# Compila a aplicação
RUN CGO_ENABLED=1 GOOS=linux go build -o main ./cmd/api/main.go

# Estágio final
FROM alpine:latest

WORKDIR /app

# Adiciona certificados CA e timezone
RUN apk add --no-cache ca-certificates tzdata

# Copia o binário compilado do estágio de build
COPY --from=builder /app/main .

# Copia o arquivo .env
COPY .env .

# Expõe a porta da aplicação
EXPOSE 8080

# Comando para executar a aplicação
CMD ["./main"]