# Definir variables
$BUILD_DIR = ".\bin"
$BIN_NAME = "Latte.exe"
$INSTALL_DIR = "C:\ProgramData\chocolatey\bin"
$CHOCOLATEY_URL = "https://chocolatey.org/install.ps1"
$GO_FILE = ".\windowsLatte.go"  # Ruta correcta al archivo Go

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

# Verificar que el archivo Go existe antes de compilar
if (-not (Test-Path -Path $GO_FILE)) {
    Write-Host "El archivo $GO_FILE no existe. Abortando la instalación."
    exit 1
}

# Compilar el programa Go
Write-Host "Compiling the Go program..."
go build -o "$BUILD_DIR\$BIN_NAME" "$GO_FILE"

# Verificar si la compilación fue exitosa
if (-not (Test-Path "$BUILD_DIR\$BIN_NAME")) {
    Write-Host "La compilación falló. Abortando la instalación."
    exit 1
}

# Instalar el binario en la ubicación de Chocolatey
Write-Host "Installing the binary to $INSTALL_DIR..."
Copy-Item -Path "$BUILD_DIR\$BIN_NAME" -Destination $INSTALL_DIR -Force

# Verificar si la instalación fue exitosa
if (Test-Path "$INSTALL_DIR\$BIN_NAME") {
    Write-Host "Installation successful! You can now use '$BIN_NAME' from anywhere."
} else {
    Write-Host "Installation failed."
}
