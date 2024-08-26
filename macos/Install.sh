#!/bin/bash

BUILD_DIR="./bin"
BIN_NAME="Latte"
INSTALL_DIR="/usr/local/bin"
GO_VERSION="1.22.6"
GO_TAR="go$GO_VERSION.darwin-amd64.tar.gz"
GO_URL="https://golang.org/dl/$GO_TAR"

check_go() {
  if command -v go &> /dev/null; then
    echo "Go ya está instalado."
    return 0
  else
    echo "Go no encontrado. Procediendo con la instalación de Go $GO_VERSION..."
    return 1
  fi
}

install_go() {
  curl -LO $GO_URL
  sudo tar -C /usr/local -xzf $GO_TAR
  echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.zshrc
  source ~/.zshrc
  rm $GO_TAR
}

mkdir -p $BUILD_DIR

if ! check_go; then
  install_go
fi

echo "Compilando el programa Go..."
go build -o $BUILD_DIR/$BIN_NAME

echo "Instalando el binario en $INSTALL_DIR..."
sudo cp $BUILD_DIR/$BIN_NAME $INSTALL_DIR/

if [ -f "$INSTALL_DIR/$BIN_NAME" ]; then
    echo "¡Instalación exitosa! Ahora puedes usar '$BIN_NAME' desde cualquier lugar."
else
    echo "La instalación falló."
fi
