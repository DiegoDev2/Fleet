# Variables
$REPO_URL = "https://github.com/CodeDiego15/LattePkg"
$CLONE_DIR = "LattePkg"
$INSTALL_SCRIPT = "Install.ps1"
$WINDOWS_DIR = "$CLONE_DIR\Windows"

# Clonar el repositorio de GitHub
Write-Host "Clonando el repositorio desde $REPO_URL..."
git clone $REPO_URL

# Comprobar si el repositorio se ha clonado con éxito
if (Test-Path $CLONE_DIR) {
    Write-Host "Repositorio clonado con éxito en $CLONE_DIR"
} else {
    Write-Host "Error al clonar el repositorio."
    exit 1
}

# Navegar a la carpeta Windows dentro del repositorio
Set-Location $WINDOWS_DIR

# Comprobar si el script de instalación existe
if (Test-Path $INSTALL_SCRIPT) {
    # Ejecutar el script de instalación
    Write-Host "Ejecutando el script de instalación..."
    & .\$INSTALL_SCRIPT

    # Volver a la carpeta original y eliminar el repositorio clonado
    Set-Location ..
    Remove-Item -Recurse -Force $CLONE_DIR
} else {
    Write-Host "No se encontró el script de instalación '$INSTALL_SCRIPT'."
    exit 1
}
