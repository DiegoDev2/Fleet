package main

// Sloth - Prometheus SLO generator
// Homepage: https://sloth.dev/

import (
	"fmt"
	
	"os/exec"
)

func installSloth() {
	// Método 1: Descargar y extraer .tar.gz
	sloth_tar_url := "https://github.com/slok/sloth/archive/refs/tags/v0.11.0.tar.gz"
	sloth_cmd_tar := exec.Command("curl", "-L", sloth_tar_url, "-o", "package.tar.gz")
	err := sloth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sloth_zip_url := "https://github.com/slok/sloth/archive/refs/tags/v0.11.0.zip"
	sloth_cmd_zip := exec.Command("curl", "-L", sloth_zip_url, "-o", "package.zip")
	err = sloth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sloth_bin_url := "https://github.com/slok/sloth/archive/refs/tags/v0.11.0.bin"
	sloth_cmd_bin := exec.Command("curl", "-L", sloth_bin_url, "-o", "binary.bin")
	err = sloth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sloth_src_url := "https://github.com/slok/sloth/archive/refs/tags/v0.11.0.src.tar.gz"
	sloth_cmd_src := exec.Command("curl", "-L", sloth_src_url, "-o", "source.tar.gz")
	err = sloth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sloth_cmd_direct := exec.Command("./binary")
	err = sloth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
