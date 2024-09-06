package main

// Favirecon - Uses favicon.ico to improve the target recon phase
// Homepage: https://github.com/edoardottt/favirecon

import (
	"fmt"
	
	"os/exec"
)

func installFavirecon() {
	// Método 1: Descargar y extraer .tar.gz
	favirecon_tar_url := "https://github.com/edoardottt/favirecon/archive/refs/tags/v0.1.2.tar.gz"
	favirecon_cmd_tar := exec.Command("curl", "-L", favirecon_tar_url, "-o", "package.tar.gz")
	err := favirecon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	favirecon_zip_url := "https://github.com/edoardottt/favirecon/archive/refs/tags/v0.1.2.zip"
	favirecon_cmd_zip := exec.Command("curl", "-L", favirecon_zip_url, "-o", "package.zip")
	err = favirecon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	favirecon_bin_url := "https://github.com/edoardottt/favirecon/archive/refs/tags/v0.1.2.bin"
	favirecon_cmd_bin := exec.Command("curl", "-L", favirecon_bin_url, "-o", "binary.bin")
	err = favirecon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	favirecon_src_url := "https://github.com/edoardottt/favirecon/archive/refs/tags/v0.1.2.src.tar.gz"
	favirecon_cmd_src := exec.Command("curl", "-L", favirecon_src_url, "-o", "source.tar.gz")
	err = favirecon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	favirecon_cmd_direct := exec.Command("./binary")
	err = favirecon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
