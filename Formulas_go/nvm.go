package main

// Nvm - Manage multiple Node.js versions
// Homepage: https://github.com/nvm-sh/nvm

import (
	"fmt"
	
	"os/exec"
)

func installNvm() {
	// Método 1: Descargar y extraer .tar.gz
	nvm_tar_url := "https://github.com/nvm-sh/nvm/archive/refs/tags/v0.40.1.tar.gz"
	nvm_cmd_tar := exec.Command("curl", "-L", nvm_tar_url, "-o", "package.tar.gz")
	err := nvm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nvm_zip_url := "https://github.com/nvm-sh/nvm/archive/refs/tags/v0.40.1.zip"
	nvm_cmd_zip := exec.Command("curl", "-L", nvm_zip_url, "-o", "package.zip")
	err = nvm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nvm_bin_url := "https://github.com/nvm-sh/nvm/archive/refs/tags/v0.40.1.bin"
	nvm_cmd_bin := exec.Command("curl", "-L", nvm_bin_url, "-o", "binary.bin")
	err = nvm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nvm_src_url := "https://github.com/nvm-sh/nvm/archive/refs/tags/v0.40.1.src.tar.gz"
	nvm_cmd_src := exec.Command("curl", "-L", nvm_src_url, "-o", "source.tar.gz")
	err = nvm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nvm_cmd_direct := exec.Command("./binary")
	err = nvm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
