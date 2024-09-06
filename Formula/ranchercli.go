package main

// RancherCli - Unified tool to manage your Rancher server
// Homepage: https://github.com/rancher/cli

import (
	"fmt"
	
	"os/exec"
)

func installRancherCli() {
	// Método 1: Descargar y extraer .tar.gz
	ranchercli_tar_url := "https://github.com/rancher/cli/archive/refs/tags/v2.9.0.tar.gz"
	ranchercli_cmd_tar := exec.Command("curl", "-L", ranchercli_tar_url, "-o", "package.tar.gz")
	err := ranchercli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ranchercli_zip_url := "https://github.com/rancher/cli/archive/refs/tags/v2.9.0.zip"
	ranchercli_cmd_zip := exec.Command("curl", "-L", ranchercli_zip_url, "-o", "package.zip")
	err = ranchercli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ranchercli_bin_url := "https://github.com/rancher/cli/archive/refs/tags/v2.9.0.bin"
	ranchercli_cmd_bin := exec.Command("curl", "-L", ranchercli_bin_url, "-o", "binary.bin")
	err = ranchercli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ranchercli_src_url := "https://github.com/rancher/cli/archive/refs/tags/v2.9.0.src.tar.gz"
	ranchercli_cmd_src := exec.Command("curl", "-L", ranchercli_src_url, "-o", "source.tar.gz")
	err = ranchercli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ranchercli_cmd_direct := exec.Command("./binary")
	err = ranchercli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
