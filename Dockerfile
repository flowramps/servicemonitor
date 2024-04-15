# Usar a imagem base do Golang para compilar a aplicação
FROM golang:latest AS builder

# Definir o diretório de trabalho dentro do contêiner
WORKDIR /app

# Copiar o arquivo de manifesto de módulos Go
COPY go.mod ./

# Baixar as dependências do módulo Go
RUN go mod download

# Copiar o código fonte da aplicação
COPY . .

# Compilar o código fonte da aplicação
RUN CGO_ENABLED=0 go build -o bifrost

# Usar uma imagem mínima do Alpine como imagem base final
FROM alpine:latest

# Copiar o executável compilado da etapa anterior para o contêiner final
COPY --from=builder /app/bifrost /app/bifrost

# Expor a porta da aplicação (se necessário)
EXPOSE 8080

# Comando para executar a aplicação quando o contêiner for iniciado
CMD ["/app/bifrost"]
