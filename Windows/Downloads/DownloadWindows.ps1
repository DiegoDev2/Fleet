# Variables
$REPO_URL = "https://github.com/CodeDiego15/LattePkg"
$CLONE_DIR = "LattePkg"

# Clonar el repositorio de GitHub
Write-Host "Clonando el repositorio desde $REPO_URL..."
git clone $REPO_URL


if (Test-Path $CLONE_DIR) {
    Write-Host "Repositorio clonado con éxito en $CLONE_DIR"
} else {
    Write-Host "Error al clonar el repositorio."
    exit 1
}

Set-Location $CLONE_DIR/Windows/Downloads/


if (Test-Path "Install.ps1") {

    Write-Host "Ejecutando el script de instalación..."
    .\Install.ps1

    Set-Location ..
    Remove-Item -Recurse -Force $CLONE_DIR
} else {
    Write-Host "No se encontró el script de instalación 'Install.ps1'."
    exit 1
}
