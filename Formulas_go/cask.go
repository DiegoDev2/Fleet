package main

// Cask - Emacs dependency management
// Homepage: https://cask.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installCask() {
	// Método 1: Descargar y extraer .tar.gz
	cask_tar_url := "https://github.com/cask/cask/archive/refs/tags/v0.9.0.tar.gz"
	cask_cmd_tar := exec.Command("curl", "-L", cask_tar_url, "-o", "package.tar.gz")
	err := cask_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cask_zip_url := "https://github.com/cask/cask/archive/refs/tags/v0.9.0.zip"
	cask_cmd_zip := exec.Command("curl", "-L", cask_zip_url, "-o", "package.zip")
	err = cask_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cask_bin_url := "https://github.com/cask/cask/archive/refs/tags/v0.9.0.bin"
	cask_cmd_bin := exec.Command("curl", "-L", cask_bin_url, "-o", "binary.bin")
	err = cask_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cask_src_url := "https://github.com/cask/cask/archive/refs/tags/v0.9.0.src.tar.gz"
	cask_cmd_src := exec.Command("curl", "-L", cask_src_url, "-o", "source.tar.gz")
	err = cask_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cask_cmd_direct := exec.Command("./binary")
	err = cask_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: coreutils")
exec.Command("latte", "install", "coreutils")
	fmt.Println("Instalando dependencia: emacs")
exec.Command("latte", "install", "emacs")
}
