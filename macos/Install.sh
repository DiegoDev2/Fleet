#!/bin/bash


BUILD_DIR="./bin"
BIN_NAME="Latte"
INSTALL_DIR="/usr/local/bin"
BREW_URL="https://brew.sh/"




check_brew() {
  if command -v brew &> /dev/null; then
    echo "Homebrew ya está instalado."
    return 0
  else
    echo "Homebrew no encontrado. Procediendo con la instalación..."
    return 1
  fi
}


install_brew() {
  
  /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
  

  if [ -f /usr/local/bin/brew ]; then
    echo 'export PATH="/usr/local/bin:$PATH"' >> ~/.zshrc
    source ~/.zshrc
  fi
}


mkdir -p $BUILD_DIR


if ! check_brew; then
  install_brew
fi


echo "Compiling the Go program..."
go build -o $BUILD_DIR/$BIN_NAME


echo "Installing the binary to $INSTALL_DIR..."
sudo cp $BUILD_DIR/$BIN_NAME $INSTALL_DIR/


if [ -f "$INSTALL_DIR/$BIN_NAME" ]; then
    echo "Installation successful! You can now use '$BIN_NAME' from anywhere."
else
    echo "Installation failed."
fi
