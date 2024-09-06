package main

// Rsyncy - Status/progress bar for rsync
// Homepage: https://github.com/laktak/rsyncy

import (
	"fmt"
	
	"os/exec"
)

func installRsyncy() {
	// Método 1: Descargar y extraer .tar.gz
	rsyncy_tar_url := "https://github.com/laktak/rsyncy/archive/refs/tags/v0.2.0-1.tar.gz"
	rsyncy_cmd_tar := exec.Command("curl", "-L", rsyncy_tar_url, "-o", "package.tar.gz")
	err := rsyncy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rsyncy_zip_url := "https://github.com/laktak/rsyncy/archive/refs/tags/v0.2.0-1.zip"
	rsyncy_cmd_zip := exec.Command("curl", "-L", rsyncy_zip_url, "-o", "package.zip")
	err = rsyncy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rsyncy_bin_url := "https://github.com/laktak/rsyncy/archive/refs/tags/v0.2.0-1.bin"
	rsyncy_cmd_bin := exec.Command("curl", "-L", rsyncy_bin_url, "-o", "binary.bin")
	err = rsyncy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rsyncy_src_url := "https://github.com/laktak/rsyncy/archive/refs/tags/v0.2.0-1.src.tar.gz"
	rsyncy_cmd_src := exec.Command("curl", "-L", rsyncy_src_url, "-o", "source.tar.gz")
	err = rsyncy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rsyncy_cmd_direct := exec.Command("./binary")
	err = rsyncy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: rsync")
exec.Command("latte", "install", "rsync")
}
