package main

// Marksman - Language Server Protocol for Markdown
// Homepage: https://github.com/artempyanykh/marksman

import (
	"fmt"
	
	"os/exec"
)

func installMarksman() {
	// Método 1: Descargar y extraer .tar.gz
	marksman_tar_url := "https://github.com/artempyanykh/marksman/archive/refs/tags/2023-12-09.tar.gz"
	marksman_cmd_tar := exec.Command("curl", "-L", marksman_tar_url, "-o", "package.tar.gz")
	err := marksman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	marksman_zip_url := "https://github.com/artempyanykh/marksman/archive/refs/tags/2023-12-09.zip"
	marksman_cmd_zip := exec.Command("curl", "-L", marksman_zip_url, "-o", "package.zip")
	err = marksman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	marksman_bin_url := "https://github.com/artempyanykh/marksman/archive/refs/tags/2023-12-09.bin"
	marksman_cmd_bin := exec.Command("curl", "-L", marksman_bin_url, "-o", "binary.bin")
	err = marksman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	marksman_src_url := "https://github.com/artempyanykh/marksman/archive/refs/tags/2023-12-09.src.tar.gz"
	marksman_cmd_src := exec.Command("curl", "-L", marksman_src_url, "-o", "source.tar.gz")
	err = marksman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	marksman_cmd_direct := exec.Command("./binary")
	err = marksman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: dotnet")
exec.Command("latte", "install", "dotnet")
}
