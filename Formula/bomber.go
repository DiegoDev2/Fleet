package main

// Bomber - Scans Software Bill of Materials for security vulnerabilities
// Homepage: https://github.com/devops-kung-fu/bomber

import (
	"fmt"
	
	"os/exec"
)

func installBomber() {
	// Método 1: Descargar y extraer .tar.gz
	bomber_tar_url := "https://github.com/devops-kung-fu/bomber/archive/refs/tags/v0.5.0.tar.gz"
	bomber_cmd_tar := exec.Command("curl", "-L", bomber_tar_url, "-o", "package.tar.gz")
	err := bomber_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bomber_zip_url := "https://github.com/devops-kung-fu/bomber/archive/refs/tags/v0.5.0.zip"
	bomber_cmd_zip := exec.Command("curl", "-L", bomber_zip_url, "-o", "package.zip")
	err = bomber_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bomber_bin_url := "https://github.com/devops-kung-fu/bomber/archive/refs/tags/v0.5.0.bin"
	bomber_cmd_bin := exec.Command("curl", "-L", bomber_bin_url, "-o", "binary.bin")
	err = bomber_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bomber_src_url := "https://github.com/devops-kung-fu/bomber/archive/refs/tags/v0.5.0.src.tar.gz"
	bomber_cmd_src := exec.Command("curl", "-L", bomber_src_url, "-o", "source.tar.gz")
	err = bomber_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bomber_cmd_direct := exec.Command("./binary")
	err = bomber_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
