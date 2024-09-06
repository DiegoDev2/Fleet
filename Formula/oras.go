package main

// Oras - OCI Registry As Storage
// Homepage: https://github.com/oras-project/oras

import (
	"fmt"
	
	"os/exec"
)

func installOras() {
	// Método 1: Descargar y extraer .tar.gz
	oras_tar_url := "https://github.com/oras-project/oras/archive/refs/tags/v1.2.0.tar.gz"
	oras_cmd_tar := exec.Command("curl", "-L", oras_tar_url, "-o", "package.tar.gz")
	err := oras_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oras_zip_url := "https://github.com/oras-project/oras/archive/refs/tags/v1.2.0.zip"
	oras_cmd_zip := exec.Command("curl", "-L", oras_zip_url, "-o", "package.zip")
	err = oras_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oras_bin_url := "https://github.com/oras-project/oras/archive/refs/tags/v1.2.0.bin"
	oras_cmd_bin := exec.Command("curl", "-L", oras_bin_url, "-o", "binary.bin")
	err = oras_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oras_src_url := "https://github.com/oras-project/oras/archive/refs/tags/v1.2.0.src.tar.gz"
	oras_cmd_src := exec.Command("curl", "-L", oras_src_url, "-o", "source.tar.gz")
	err = oras_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oras_cmd_direct := exec.Command("./binary")
	err = oras_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
