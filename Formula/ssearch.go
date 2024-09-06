package main

// SSearch - Web search from the terminal
// Homepage: https://github.com/zquestz/s

import (
	"fmt"
	
	"os/exec"
)

func installSSearch() {
	// Método 1: Descargar y extraer .tar.gz
	ssearch_tar_url := "https://github.com/zquestz/s/archive/refs/tags/v0.7.1.tar.gz"
	ssearch_cmd_tar := exec.Command("curl", "-L", ssearch_tar_url, "-o", "package.tar.gz")
	err := ssearch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ssearch_zip_url := "https://github.com/zquestz/s/archive/refs/tags/v0.7.1.zip"
	ssearch_cmd_zip := exec.Command("curl", "-L", ssearch_zip_url, "-o", "package.zip")
	err = ssearch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ssearch_bin_url := "https://github.com/zquestz/s/archive/refs/tags/v0.7.1.bin"
	ssearch_cmd_bin := exec.Command("curl", "-L", ssearch_bin_url, "-o", "binary.bin")
	err = ssearch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ssearch_src_url := "https://github.com/zquestz/s/archive/refs/tags/v0.7.1.src.tar.gz"
	ssearch_cmd_src := exec.Command("curl", "-L", ssearch_src_url, "-o", "source.tar.gz")
	err = ssearch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ssearch_cmd_direct := exec.Command("./binary")
	err = ssearch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
