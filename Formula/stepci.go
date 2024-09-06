package main

// Stepci - API Testing and Monitoring made simple
// Homepage: https://stepci.com

import (
	"fmt"
	
	"os/exec"
)

func installStepci() {
	// Método 1: Descargar y extraer .tar.gz
	stepci_tar_url := "https://registry.npmjs.org/stepci/-/stepci-2.8.2.tgz"
	stepci_cmd_tar := exec.Command("curl", "-L", stepci_tar_url, "-o", "package.tar.gz")
	err := stepci_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stepci_zip_url := "https://registry.npmjs.org/stepci/-/stepci-2.8.2.tgz"
	stepci_cmd_zip := exec.Command("curl", "-L", stepci_zip_url, "-o", "package.zip")
	err = stepci_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stepci_bin_url := "https://registry.npmjs.org/stepci/-/stepci-2.8.2.tgz"
	stepci_cmd_bin := exec.Command("curl", "-L", stepci_bin_url, "-o", "binary.bin")
	err = stepci_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stepci_src_url := "https://registry.npmjs.org/stepci/-/stepci-2.8.2.tgz"
	stepci_cmd_src := exec.Command("curl", "-L", stepci_src_url, "-o", "source.tar.gz")
	err = stepci_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stepci_cmd_direct := exec.Command("./binary")
	err = stepci_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
