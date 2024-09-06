package main

// Slacknimate - Text animation for Slack messages
// Homepage: https://github.com/mroth/slacknimate

import (
	"fmt"
	
	"os/exec"
)

func installSlacknimate() {
	// Método 1: Descargar y extraer .tar.gz
	slacknimate_tar_url := "https://github.com/mroth/slacknimate/archive/refs/tags/v1.1.0.tar.gz"
	slacknimate_cmd_tar := exec.Command("curl", "-L", slacknimate_tar_url, "-o", "package.tar.gz")
	err := slacknimate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	slacknimate_zip_url := "https://github.com/mroth/slacknimate/archive/refs/tags/v1.1.0.zip"
	slacknimate_cmd_zip := exec.Command("curl", "-L", slacknimate_zip_url, "-o", "package.zip")
	err = slacknimate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	slacknimate_bin_url := "https://github.com/mroth/slacknimate/archive/refs/tags/v1.1.0.bin"
	slacknimate_cmd_bin := exec.Command("curl", "-L", slacknimate_bin_url, "-o", "binary.bin")
	err = slacknimate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	slacknimate_src_url := "https://github.com/mroth/slacknimate/archive/refs/tags/v1.1.0.src.tar.gz"
	slacknimate_cmd_src := exec.Command("curl", "-L", slacknimate_src_url, "-o", "source.tar.gz")
	err = slacknimate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	slacknimate_cmd_direct := exec.Command("./binary")
	err = slacknimate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
