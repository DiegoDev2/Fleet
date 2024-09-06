package main

// Wush - Transfer files between computers via WireGuard
// Homepage: https://github.com/coder/wush

import (
	"fmt"
	
	"os/exec"
)

func installWush() {
	// Método 1: Descargar y extraer .tar.gz
	wush_tar_url := "https://github.com/coder/wush/archive/refs/tags/v0.1.2.tar.gz"
	wush_cmd_tar := exec.Command("curl", "-L", wush_tar_url, "-o", "package.tar.gz")
	err := wush_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wush_zip_url := "https://github.com/coder/wush/archive/refs/tags/v0.1.2.zip"
	wush_cmd_zip := exec.Command("curl", "-L", wush_zip_url, "-o", "package.zip")
	err = wush_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wush_bin_url := "https://github.com/coder/wush/archive/refs/tags/v0.1.2.bin"
	wush_cmd_bin := exec.Command("curl", "-L", wush_bin_url, "-o", "binary.bin")
	err = wush_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wush_src_url := "https://github.com/coder/wush/archive/refs/tags/v0.1.2.src.tar.gz"
	wush_cmd_src := exec.Command("curl", "-L", wush_src_url, "-o", "source.tar.gz")
	err = wush_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wush_cmd_direct := exec.Command("./binary")
	err = wush_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
