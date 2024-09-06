package main

// Minder - CLI for interacting with Stacklok's Minder platform
// Homepage: https://minder-docs.stacklok.dev

import (
	"fmt"
	
	"os/exec"
)

func installMinder() {
	// Método 1: Descargar y extraer .tar.gz
	minder_tar_url := "https://github.com/stacklok/minder/archive/refs/tags/v0.0.62.tar.gz"
	minder_cmd_tar := exec.Command("curl", "-L", minder_tar_url, "-o", "package.tar.gz")
	err := minder_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minder_zip_url := "https://github.com/stacklok/minder/archive/refs/tags/v0.0.62.zip"
	minder_cmd_zip := exec.Command("curl", "-L", minder_zip_url, "-o", "package.zip")
	err = minder_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minder_bin_url := "https://github.com/stacklok/minder/archive/refs/tags/v0.0.62.bin"
	minder_cmd_bin := exec.Command("curl", "-L", minder_bin_url, "-o", "binary.bin")
	err = minder_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minder_src_url := "https://github.com/stacklok/minder/archive/refs/tags/v0.0.62.src.tar.gz"
	minder_cmd_src := exec.Command("curl", "-L", minder_src_url, "-o", "source.tar.gz")
	err = minder_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minder_cmd_direct := exec.Command("./binary")
	err = minder_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
