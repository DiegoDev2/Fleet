#!/bin/bash

REPO_URL="https://github.com/CodeDiego15/LattePkg"
CLONE_DIR="LattePkg"

echo "Clonando el repositorio desde $REPO_URL..."
git clone $REPO_URL

if [ -d "$CLONE_DIR" ]; then
    echo "Repositorio clonado con éxito en $CLONE_DIR"
else
    echo "Error al clonar el repositorio."
    exit 1
fi

cd $CLONE_DIR

if [ -f "macos/Install.sh" ]; then
    chmod +x macos/Install.sh
    echo "Ejecutando el script de instalación..."
    ./macos/Install.sh
    cd ..
    rm -rf $CLONE_DIR
else
    echo "No se encontró el script de instalación 'Install.sh'."
    exit 1
fi
