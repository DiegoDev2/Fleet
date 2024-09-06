package main

// Pipgrip - Lightweight pip dependency resolver
// Homepage: https://github.com/ddelange/pipgrip

import (
	"fmt"
	
	"os/exec"
)

func installPipgrip() {
	// Método 1: Descargar y extraer .tar.gz
	pipgrip_tar_url := "https://files.pythonhosted.org/packages/fb/dc/f89b89e72e541bb5ffa25cbaf1f9c92d2c2187644c8972772aafb7bf0009/pipgrip-0.10.13.tar.gz"
	pipgrip_cmd_tar := exec.Command("curl", "-L", pipgrip_tar_url, "-o", "package.tar.gz")
	err := pipgrip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pipgrip_zip_url := "https://files.pythonhosted.org/packages/fb/dc/f89b89e72e541bb5ffa25cbaf1f9c92d2c2187644c8972772aafb7bf0009/pipgrip-0.10.13.zip"
	pipgrip_cmd_zip := exec.Command("curl", "-L", pipgrip_zip_url, "-o", "package.zip")
	err = pipgrip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pipgrip_bin_url := "https://files.pythonhosted.org/packages/fb/dc/f89b89e72e541bb5ffa25cbaf1f9c92d2c2187644c8972772aafb7bf0009/pipgrip-0.10.13.bin"
	pipgrip_cmd_bin := exec.Command("curl", "-L", pipgrip_bin_url, "-o", "binary.bin")
	err = pipgrip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pipgrip_src_url := "https://files.pythonhosted.org/packages/fb/dc/f89b89e72e541bb5ffa25cbaf1f9c92d2c2187644c8972772aafb7bf0009/pipgrip-0.10.13.src.tar.gz"
	pipgrip_cmd_src := exec.Command("curl", "-L", pipgrip_src_url, "-o", "source.tar.gz")
	err = pipgrip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pipgrip_cmd_direct := exec.Command("./binary")
	err = pipgrip_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
