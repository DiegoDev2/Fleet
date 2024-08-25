
$REPO_URL = "https://github.com/CodeDiego15/LattePkg"
$CLONE_DIR = "LattePkg"
$INSTALL_SCRIPT = "Install.ps1"
$WINDOWS_DIR = "$CLONE_DIR\Windows"


Write-Host "Clonando el repositorio desde $REPO_URL..."
git clone $REPO_URL

if (Test-Path $CLONE_DIR) {
    Write-Host "Repositorio clonado con éxito en $CLONE_DIR"
} else {
    Write-Host "Error al clonar el repositorio."
    exit 1
}


Set-Location $WINDOWS_DIR


if (Test-Path $INSTALL_SCRIPT) {
    
    Write-Host "Ejecutando el script de instalación..."
    & .\$INSTALL_SCRIPT

   
    Set-Location ..
    Remove-Item -Recurse -Force $CLONE_DIR
} else {
    Write-Host "No se encontró el script de instalación '$INSTALL_SCRIPT'."
    exit 1
}
