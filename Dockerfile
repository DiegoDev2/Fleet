# Usar una imagen base de Alpine Linux
FROM alpine:3.16

# Establecer el directorio de trabajo dentro del contenedor
WORKDIR /app

# Copiar todas las carpetas de fórmulas dentro del contenedor
COPY Formula /app/Formula

# Comando por defecto al iniciar el contenedor
CMD ["ls", "/app/Formula"]

