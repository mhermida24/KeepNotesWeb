FROM golang:1.25.1

WORKDIR /app

# Instalar Air
RUN go install github.com/air-verse/air@latest

# Copiar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Copiar el código
COPY . .

# Crear carpeta tmp para el binario
RUN mkdir -p tmp

# Exponer puerto
EXPOSE 8080

# Iniciar Air
CMD ["air", "-c", ".air.toml"]
