package main

// Dependabot - Tool for testing and debugging Dependabot update jobs
// Homepage: https://github.com/dependabot/cli

import (
	"fmt"
	
	"os/exec"
)

func installDependabot() {
	// Método 1: Descargar y extraer .tar.gz
	dependabot_tar_url := "https://github.com/dependabot/cli/archive/refs/tags/v1.54.0.tar.gz"
	dependabot_cmd_tar := exec.Command("curl", "-L", dependabot_tar_url, "-o", "package.tar.gz")
	err := dependabot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dependabot_zip_url := "https://github.com/dependabot/cli/archive/refs/tags/v1.54.0.zip"
	dependabot_cmd_zip := exec.Command("curl", "-L", dependabot_zip_url, "-o", "package.zip")
	err = dependabot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dependabot_bin_url := "https://github.com/dependabot/cli/archive/refs/tags/v1.54.0.bin"
	dependabot_cmd_bin := exec.Command("curl", "-L", dependabot_bin_url, "-o", "binary.bin")
	err = dependabot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dependabot_src_url := "https://github.com/dependabot/cli/archive/refs/tags/v1.54.0.src.tar.gz"
	dependabot_cmd_src := exec.Command("curl", "-L", dependabot_src_url, "-o", "source.tar.gz")
	err = dependabot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dependabot_cmd_direct := exec.Command("./binary")
	err = dependabot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
