#!/bin/bash

# Definir variables
BUILD_DIR="./bin"
BIN_NAME="Latte"
INSTALL_DIR="/usr/local/bin"
BREW_URL="https://brew.sh/"



# Función para comprobar si Homebrew está instalado
check_brew() {
  if command -v brew &> /dev/null; then
    echo "Homebrew ya está instalado."
    return 0
  else
    echo "Homebrew no encontrado. Procediendo con la instalación..."
    return 1
  fi
}

# Función para instalar Homebrew
install_brew() {
  # Instalar Homebrew
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  
  # Añadir Homebrew al PATH (para zsh)
  if [ -f /usr/local/bin/brew ]; then
    echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.zshrc
    source ~/.zshrc
  fi
}

# Crear el directorio de compilación si no existe
mkdir -p $BUILD_DIR

# Comprobar e instalar Homebrew si es necesario
if ! check_brew; then
  install_brew
fi

# Compilar el programa
echo "Compiling the Go program..."
go build -o $BUILD_DIR/$BIN_NAME

# Instalar el binario
echo "Installing the binary to $INSTALL_DIR..."
sudo cp $BUILD_DIR/$BIN_NAME $INSTALL_DIR/

# Verificar la instalación
if [ -f "$INSTALL_DIR/$BIN_NAME" ]; then
    echo "Installation successful! You can now use '$BIN_NAME' from anywhere."
else
    echo "Installation failed."
fi
