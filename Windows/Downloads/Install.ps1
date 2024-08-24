# Definir variables
$BUILD_DIR = ".\bin"
$BIN_NAME = "Latte.exe"
$INSTALL_DIR = "C:\ProgramData\chocolatey\bin"
$CHOCOLATEY_URL = "https://chocolatey.org/install.ps1"

# Función para comprobar si Chocolatey está instalado
function Check-Chocolatey {
    if (Get-Command choco -ErrorAction SilentlyContinue) {
        Write-Host "Chocolatey ya está instalado."
        return $true
    } else {
        Write-Host "Chocolatey no encontrado. Procediendo con la instalación..."
        return $false
    }
}

# Función para instalar Chocolatey
function Install-Chocolatey {
    # Instalar Chocolatey
    Set-ExecutionPolicy Bypass -Scope Process -Force;
    [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072;
    Invoke-Expression ((New-Object System.Net.WebClient).DownloadString($CHOCOLATEY_URL))
}

# Crear el directorio de compilación si no existe
if (-not (Test-Path -Path $BUILD_DIR)) {
    New-Item -ItemType Directory -Path $BUILD_DIR | Out-Null
}

# Comprobar e instalar Chocolatey si es necesario
if (-not (Check-Chocolatey)) {
    Install-Chocolatey
}

# Compilar el programa
Write-Host "Compiling the Go program..."
go build -o "$BUILD_DIR\$BIN_NAME"

# Instalar el binario
Write-Host "Installing the binary to $INSTALL_DIR..."
Copy-Item -Path "$BUILD_DIR\$BIN_NAME" -Destination $INSTALL_DIR -Force

# Verificar la instalación
if (Test-Path "$INSTALL_DIR\$BIN_NAME") {
    Write-Host "Installation successful! You can now use '$BIN_NAME' from anywhere."
} else {
    Write-Host "Installation failed."
}
