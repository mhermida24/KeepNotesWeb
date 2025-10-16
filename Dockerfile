# Etapa 1: Compilación
FROM golang:1.25.1

# Establecer el directorio de trabajo
WORKDIR /app

# Copiar los archivos de dependencias y descargarlas
COPY go.mod go.sum ./


RUN go mod download

# Copiar el resto del código fuente
COPY . .

# Compilar el binario
RUN go build -o main .


# Tells Docker which network port your container listens on
EXPOSE 8080
 
# Specifies the executable command that runs when the container starts
CMD [ "./main" ]
