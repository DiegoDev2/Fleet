package main

// Tart - macOS and Linux VMs on Apple Silicon to use in CI and other automations
// Homepage: https://github.com/cirruslabs/tart

import (
	"fmt"
	
	"os/exec"
)

func installTart() {
	// Método 1: Descargar y extraer .tar.gz
	tart_tar_url := "https://github.com/cirruslabs/tart/archive/refs/tags/0.38.0.tar.gz"
	tart_cmd_tar := exec.Command("curl", "-L", tart_tar_url, "-o", "package.tar.gz")
	err := tart_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tart_zip_url := "https://github.com/cirruslabs/tart/archive/refs/tags/0.38.0.zip"
	tart_cmd_zip := exec.Command("curl", "-L", tart_zip_url, "-o", "package.zip")
	err = tart_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tart_bin_url := "https://github.com/cirruslabs/tart/archive/refs/tags/0.38.0.bin"
	tart_cmd_bin := exec.Command("curl", "-L", tart_bin_url, "-o", "binary.bin")
	err = tart_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tart_src_url := "https://github.com/cirruslabs/tart/archive/refs/tags/0.38.0.src.tar.gz"
	tart_cmd_src := exec.Command("curl", "-L", tart_src_url, "-o", "source.tar.gz")
	err = tart_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tart_cmd_direct := exec.Command("./binary")
	err = tart_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
