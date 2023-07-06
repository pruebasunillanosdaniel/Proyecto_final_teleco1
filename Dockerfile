# syntax=docker/dockerfile:1
# importar docker imagen con    golang implementado en Linux -  Alpine
FROM golang:1.20-alpine as builder  
#Definir directorio  de trabajo
WORKDIR /usr/src/app
#Copiar go mod y go.sum para verificar repositorios y descargar posibles actualizaciones
COPY go.mod go.sum ./
RUN go mod download && go mod verify
#Copiar archivos de proyecto  actuales  a la imagen en el  WORKDIR
COPY . .
#Copiar archivos fuentes y crear ejecutable (binario) en la ruta /usr/local/bin/app ,con nombre app
RUN go build -v -o /usr/local/bin/app
# crear  grupo y usuario linux app para  ejecutar binario
RUN addgroup -S app && adduser -S app -G app
# crear  permisos para usuario app sobre el binario
RUN chown app /usr/local/bin/app
# cambiar a usuario app
USER app
# ejecutar binario
ENTRYPOINT ["/usr/local/bin/app"]
# habilitar puerto 9000
EXPOSE 80
