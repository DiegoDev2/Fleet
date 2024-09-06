package main

// Lci - Interpreter for the lambda calculus
// Homepage: https://www.chatzi.org/lci/

import (
	"fmt"
	
	"os/exec"
)

func installLci() {
	// Método 1: Descargar y extraer .tar.gz
	lci_tar_url := "https://github.com/chatziko/lci/releases/download/v1.0/lci-1.0.tar.gz"
	lci_cmd_tar := exec.Command("curl", "-L", lci_tar_url, "-o", "package.tar.gz")
	err := lci_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lci_zip_url := "https://github.com/chatziko/lci/releases/download/v1.0/lci-1.0.zip"
	lci_cmd_zip := exec.Command("curl", "-L", lci_zip_url, "-o", "package.zip")
	err = lci_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lci_bin_url := "https://github.com/chatziko/lci/releases/download/v1.0/lci-1.0.bin"
	lci_cmd_bin := exec.Command("curl", "-L", lci_bin_url, "-o", "binary.bin")
	err = lci_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lci_src_url := "https://github.com/chatziko/lci/releases/download/v1.0/lci-1.0.src.tar.gz"
	lci_cmd_src := exec.Command("curl", "-L", lci_src_url, "-o", "source.tar.gz")
	err = lci_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lci_cmd_direct := exec.Command("./binary")
	err = lci_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
