package main

// Aichat - All-in-one AI-Powered CLI Chat & Copilot
// Homepage: https://github.com/sigoden/aichat

import (
	"fmt"
	
	"os/exec"
)

func installAichat() {
	// Método 1: Descargar y extraer .tar.gz
	aichat_tar_url := "https://github.com/sigoden/aichat/archive/refs/tags/v0.21.1.tar.gz"
	aichat_cmd_tar := exec.Command("curl", "-L", aichat_tar_url, "-o", "package.tar.gz")
	err := aichat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aichat_zip_url := "https://github.com/sigoden/aichat/archive/refs/tags/v0.21.1.zip"
	aichat_cmd_zip := exec.Command("curl", "-L", aichat_zip_url, "-o", "package.zip")
	err = aichat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aichat_bin_url := "https://github.com/sigoden/aichat/archive/refs/tags/v0.21.1.bin"
	aichat_cmd_bin := exec.Command("curl", "-L", aichat_bin_url, "-o", "binary.bin")
	err = aichat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aichat_src_url := "https://github.com/sigoden/aichat/archive/refs/tags/v0.21.1.src.tar.gz"
	aichat_cmd_src := exec.Command("curl", "-L", aichat_src_url, "-o", "source.tar.gz")
	err = aichat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aichat_cmd_direct := exec.Command("./binary")
	err = aichat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
