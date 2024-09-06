package main

// Litmusctl - Command-line interface for interacting with LitmusChaos
// Homepage: https://litmuschaos.io

import (
	"fmt"
	
	"os/exec"
)

func installLitmusctl() {
	// Método 1: Descargar y extraer .tar.gz
	litmusctl_tar_url := "https://github.com/litmuschaos/litmusctl/archive/refs/tags/1.9.0.tar.gz"
	litmusctl_cmd_tar := exec.Command("curl", "-L", litmusctl_tar_url, "-o", "package.tar.gz")
	err := litmusctl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	litmusctl_zip_url := "https://github.com/litmuschaos/litmusctl/archive/refs/tags/1.9.0.zip"
	litmusctl_cmd_zip := exec.Command("curl", "-L", litmusctl_zip_url, "-o", "package.zip")
	err = litmusctl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	litmusctl_bin_url := "https://github.com/litmuschaos/litmusctl/archive/refs/tags/1.9.0.bin"
	litmusctl_cmd_bin := exec.Command("curl", "-L", litmusctl_bin_url, "-o", "binary.bin")
	err = litmusctl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	litmusctl_src_url := "https://github.com/litmuschaos/litmusctl/archive/refs/tags/1.9.0.src.tar.gz"
	litmusctl_cmd_src := exec.Command("curl", "-L", litmusctl_src_url, "-o", "source.tar.gz")
	err = litmusctl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	litmusctl_cmd_direct := exec.Command("./binary")
	err = litmusctl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
