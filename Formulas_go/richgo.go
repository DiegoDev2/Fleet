package main

// Richgo - Enrich `go test` outputs with text decorations
// Homepage: https://github.com/kyoh86/richgo

import (
	"fmt"
	
	"os/exec"
)

func installRichgo() {
	// Método 1: Descargar y extraer .tar.gz
	richgo_tar_url := "https://github.com/kyoh86/richgo/archive/refs/tags/v0.3.12.tar.gz"
	richgo_cmd_tar := exec.Command("curl", "-L", richgo_tar_url, "-o", "package.tar.gz")
	err := richgo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	richgo_zip_url := "https://github.com/kyoh86/richgo/archive/refs/tags/v0.3.12.zip"
	richgo_cmd_zip := exec.Command("curl", "-L", richgo_zip_url, "-o", "package.zip")
	err = richgo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	richgo_bin_url := "https://github.com/kyoh86/richgo/archive/refs/tags/v0.3.12.bin"
	richgo_cmd_bin := exec.Command("curl", "-L", richgo_bin_url, "-o", "binary.bin")
	err = richgo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	richgo_src_url := "https://github.com/kyoh86/richgo/archive/refs/tags/v0.3.12.src.tar.gz"
	richgo_cmd_src := exec.Command("curl", "-L", richgo_src_url, "-o", "source.tar.gz")
	err = richgo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	richgo_cmd_direct := exec.Command("./binary")
	err = richgo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
