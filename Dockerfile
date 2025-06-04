# Etapa de construcción
FROM golang:1.20 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o cabys-api ./cmd/main.go

# Etapa de ejecución
FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/cabys-api .
COPY data/cabys.json ./data/cabys.json
EXPOSE 8080
CMD ["./cabys-api"]