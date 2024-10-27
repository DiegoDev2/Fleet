#!/bin/bash

# Colores para el output
BOLD_GREEN='\033[1;32m'
BOLD_RED='\033[1;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Verificar si Go está instalado
if ! command -v go &> /dev/null; then
    echo -e "${YELLOW}Go is not installed. Downloading and installing Go...${NC}"
    
    # Determinar si la Mac es ARM (Apple Silicon) o Intel
    ARCH=$(uname -m)
    
    if [ "$ARCH" = "arm64" ]; then
        echo -e "${YELLOW}Detected Apple Silicon (M1/M2)...${NC}"
        GO_VERSION="1.20.4"
        GO_TAR="go$GO_VERSION.darwin-arm64.tar.gz"
    else
        echo -e "${YELLOW}Detected Intel...${NC}"
        GO_VERSION="1.20.4"
        GO_TAR="go$GO_VERSION.darwin-amd64.tar.gz"
    fi
    
    # Descargar Go
    curl -LO "https://golang.org/dl/$GO_TAR"
    
    # Descomprimir y mover Go a /usr/local
    sudo tar -C /usr/local -xzf "$GO_TAR"
    
    # Añadir Go al PATH
    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bash_profile
    source ~/.bash_profile

    # Verificar si Go se instaló correctamente
    if ! command -v go &> /dev/null; then
        echo -e "${BOLD_RED}Error: Failed to install Go.${NC}"
        exit 1
    fi
else
    echo -e "${BOLD_GREEN}Go is already installed.${NC}"
fi

# Clonar el repositorio de Fleet
echo -e "${YELLOW}Cloning the Fleet repository...${NC}"
git clone https://github.com/DiegoDev2/Fleet
cd Fleet || { echo -e "${BOLD_RED}Error: Failed to enter the Fleet directory.${NC}"; exit 1; }

# Compilar el proyecto en Go
echo -e "${YELLOW}Building Fleet...${NC}"
go build -o fleet

# Mover el binario a /usr/local/bin
echo -e "${YELLOW}Moving Fleet to /usr/local/bin...${NC}"
sudo mv fleet /usr/local/bin

# Verificar si la instalación fue exitosa
if command -v fleet &> /dev/null; then
    echo -e "${BOLD_GREEN}Fleet installed successfully! 🎉${NC}"
else
    echo -e "${BOLD_RED}Error: Fleet installation failed.${NC}"
fi
