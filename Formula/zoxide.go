package main

// Zoxide - Shell extension to navigate your filesystem faster
// Homepage: https://github.com/ajeetdsouza/zoxide

import (
	"fmt"
	
	"os/exec"
)

func installZoxide() {
	// Método 1: Descargar y extraer .tar.gz
	zoxide_tar_url := "https://github.com/ajeetdsouza/zoxide/archive/refs/tags/v0.9.4.tar.gz"
	zoxide_cmd_tar := exec.Command("curl", "-L", zoxide_tar_url, "-o", "package.tar.gz")
	err := zoxide_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zoxide_zip_url := "https://github.com/ajeetdsouza/zoxide/archive/refs/tags/v0.9.4.zip"
	zoxide_cmd_zip := exec.Command("curl", "-L", zoxide_zip_url, "-o", "package.zip")
	err = zoxide_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zoxide_bin_url := "https://github.com/ajeetdsouza/zoxide/archive/refs/tags/v0.9.4.bin"
	zoxide_cmd_bin := exec.Command("curl", "-L", zoxide_bin_url, "-o", "binary.bin")
	err = zoxide_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zoxide_src_url := "https://github.com/ajeetdsouza/zoxide/archive/refs/tags/v0.9.4.src.tar.gz"
	zoxide_cmd_src := exec.Command("curl", "-L", zoxide_src_url, "-o", "source.tar.gz")
	err = zoxide_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zoxide_cmd_direct := exec.Command("./binary")
	err = zoxide_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
