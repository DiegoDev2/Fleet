package main

// Peco - Simplistic interactive filtering tool
// Homepage: https://github.com/peco/peco

import (
	"fmt"
	
	"os/exec"
)

func installPeco() {
	// Método 1: Descargar y extraer .tar.gz
	peco_tar_url := "https://github.com/peco/peco/archive/refs/tags/v0.5.11.tar.gz"
	peco_cmd_tar := exec.Command("curl", "-L", peco_tar_url, "-o", "package.tar.gz")
	err := peco_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	peco_zip_url := "https://github.com/peco/peco/archive/refs/tags/v0.5.11.zip"
	peco_cmd_zip := exec.Command("curl", "-L", peco_zip_url, "-o", "package.zip")
	err = peco_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	peco_bin_url := "https://github.com/peco/peco/archive/refs/tags/v0.5.11.bin"
	peco_cmd_bin := exec.Command("curl", "-L", peco_bin_url, "-o", "binary.bin")
	err = peco_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	peco_src_url := "https://github.com/peco/peco/archive/refs/tags/v0.5.11.src.tar.gz"
	peco_cmd_src := exec.Command("curl", "-L", peco_src_url, "-o", "source.tar.gz")
	err = peco_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	peco_cmd_direct := exec.Command("./binary")
	err = peco_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
