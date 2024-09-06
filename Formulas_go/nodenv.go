package main

// Nodenv - Manage multiple NodeJS versions
// Homepage: https://github.com/nodenv/nodenv

import (
	"fmt"
	
	"os/exec"
)

func installNodenv() {
	// Método 1: Descargar y extraer .tar.gz
	nodenv_tar_url := "https://github.com/nodenv/nodenv/archive/refs/tags/v1.5.0.tar.gz"
	nodenv_cmd_tar := exec.Command("curl", "-L", nodenv_tar_url, "-o", "package.tar.gz")
	err := nodenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nodenv_zip_url := "https://github.com/nodenv/nodenv/archive/refs/tags/v1.5.0.zip"
	nodenv_cmd_zip := exec.Command("curl", "-L", nodenv_zip_url, "-o", "package.zip")
	err = nodenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nodenv_bin_url := "https://github.com/nodenv/nodenv/archive/refs/tags/v1.5.0.bin"
	nodenv_cmd_bin := exec.Command("curl", "-L", nodenv_bin_url, "-o", "binary.bin")
	err = nodenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nodenv_src_url := "https://github.com/nodenv/nodenv/archive/refs/tags/v1.5.0.src.tar.gz"
	nodenv_cmd_src := exec.Command("curl", "-L", nodenv_src_url, "-o", "source.tar.gz")
	err = nodenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nodenv_cmd_direct := exec.Command("./binary")
	err = nodenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node-build")
exec.Command("latte", "install", "node-build")
}
