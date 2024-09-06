package main

// Gops - Tool to list and diagnose Go processes currently running on your system
// Homepage: https://github.com/google/gops

import (
	"fmt"
	
	"os/exec"
)

func installGops() {
	// Método 1: Descargar y extraer .tar.gz
	gops_tar_url := "https://github.com/google/gops/archive/refs/tags/v0.3.28.tar.gz"
	gops_cmd_tar := exec.Command("curl", "-L", gops_tar_url, "-o", "package.tar.gz")
	err := gops_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gops_zip_url := "https://github.com/google/gops/archive/refs/tags/v0.3.28.zip"
	gops_cmd_zip := exec.Command("curl", "-L", gops_zip_url, "-o", "package.zip")
	err = gops_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gops_bin_url := "https://github.com/google/gops/archive/refs/tags/v0.3.28.bin"
	gops_cmd_bin := exec.Command("curl", "-L", gops_bin_url, "-o", "binary.bin")
	err = gops_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gops_src_url := "https://github.com/google/gops/archive/refs/tags/v0.3.28.src.tar.gz"
	gops_cmd_src := exec.Command("curl", "-L", gops_src_url, "-o", "source.tar.gz")
	err = gops_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gops_cmd_direct := exec.Command("./binary")
	err = gops_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
