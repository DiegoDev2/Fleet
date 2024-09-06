package main

// Bvm - Version manager for all binaries
// Homepage: https://github.com/bvm/bvm

import (
	"fmt"
	
	"os/exec"
)

func installBvm() {
	// Método 1: Descargar y extraer .tar.gz
	bvm_tar_url := "https://github.com/bvm/bvm/archive/refs/tags/0.4.2.tar.gz"
	bvm_cmd_tar := exec.Command("curl", "-L", bvm_tar_url, "-o", "package.tar.gz")
	err := bvm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bvm_zip_url := "https://github.com/bvm/bvm/archive/refs/tags/0.4.2.zip"
	bvm_cmd_zip := exec.Command("curl", "-L", bvm_zip_url, "-o", "package.zip")
	err = bvm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bvm_bin_url := "https://github.com/bvm/bvm/archive/refs/tags/0.4.2.bin"
	bvm_cmd_bin := exec.Command("curl", "-L", bvm_bin_url, "-o", "binary.bin")
	err = bvm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bvm_src_url := "https://github.com/bvm/bvm/archive/refs/tags/0.4.2.src.tar.gz"
	bvm_cmd_src := exec.Command("curl", "-L", bvm_src_url, "-o", "source.tar.gz")
	err = bvm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bvm_cmd_direct := exec.Command("./binary")
	err = bvm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
