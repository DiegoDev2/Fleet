#!/bin/bash

# Variables
REPO_URL="https://github.com/CodeDiego15/LattePkg"
CLONE_DIR="LattePkg"

# Clonar el repositorio de GitHub
echo "Clonando el repositorio desde $REPO_URL..."
git clone $REPO_URL

# Comprobar si el repositorio se ha clonado con éxito
if [ -d "$CLONE_DIR" ]; then
    echo "Repositorio clonado con éxito en $CLONE_DIR"
else
    echo "Error al clonar el repositorio."
    exit 1
fi

# Navegar a la carpeta del repositorio
cd $CLONE_DIR

# Comprobar si el script de instalación existe
if [ -f "Install.sh" ]; then
    # Hacer que el script de instalación sea ejecutable
    chmod +x Install.sh

    # Ejecutar el script de instalación
    echo "Ejecutando el script de instalación..."
    ./Install.sh

    cd ..
    rm -rm $CLONE_DIR
else
    echo "No se encontró el script de instalación 'install.sh'."
    exit 1
fi
