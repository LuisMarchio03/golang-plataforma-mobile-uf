#!/bin/bash

echo "Verificando container do PostgreSQL..."
if ! docker ps | grep -q postgres; then
    echo "Iniciando container PostgreSQL..."
    docker run --name postgres-plataforma -e POSTGRES_PASSWORD=admin123 -e POSTGRES_USER=admin -e POSTGRES_DB=plataforma_mobile -p 5432:5432 -d postgres:latest
    
    # Aguarda o PostgreSQL iniciar
    echo "Aguardando PostgreSQL iniciar..."
    sleep 5
else
    echo "Container PostgreSQL já está rodando"
fi

# Verifica se o banco de dados existe
echo "Verificando banco de dados..."
if ! docker exec postgres-plataforma psql -U admin -lqt | cut -d \| -f 1 | grep -qw plataforma_mobile; then
    echo "Criando banco de dados..."
    docker exec postgres-plataforma psql -U admin -c "CREATE DATABASE plataforma_mobile;"
fi

# Executa as migrações
echo "Executando migrações..."
docker exec -i postgres-plataforma psql -U admin -d plataforma_mobile < ./internal/database/migrations/init.sql

# Cria uma rede Docker se não existir
echo "Verificando rede Docker..."
if ! docker network ls | grep -q plataforma-network; then
    echo "Criando rede Docker..."
    docker network create plataforma-network
fi

# Conecta o container do PostgreSQL à rede
echo "Conectando PostgreSQL à rede..."
docker network connect plataforma-network postgres-plataforma || true

# Build da imagem Go
echo "Construindo imagem Go..."
docker build -t plataforma-mobile-api -f Dockerfile .

# Remove container anterior se existir
echo "Removendo container Go anterior se existir..."
docker rm -f plataforma-mobile-api || true

# Inicia o container Go
echo "Iniciando aplicação Go em container..."
docker run --name plataforma-mobile-api \
    --network plataforma-network \
    -e DB_HOST=postgres-plataforma \
    -e DB_PORT=5432 \
    -e DB_USER=admin \
    -e DB_PASSWORD=admin123 \
    -e DB_NAME=plataforma_mobile \
    -e SERVER_PORT=8080 \
    -e LOG_LEVEL=INFO \
    -p 8080:8080 \
    -d plataforma-mobile-api

echo "Aplicação iniciada com sucesso!"
echo "API disponível em: http://localhost:8080"
go run ./cmd/api/main.go