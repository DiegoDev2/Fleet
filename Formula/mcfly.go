package main

// Mcfly - Fly through your shell history
// Homepage: https://github.com/cantino/mcfly

import (
	"fmt"
	
	"os/exec"
)

func installMcfly() {
	// Método 1: Descargar y extraer .tar.gz
	mcfly_tar_url := "https://github.com/cantino/mcfly/archive/refs/tags/v0.9.2.tar.gz"
	mcfly_cmd_tar := exec.Command("curl", "-L", mcfly_tar_url, "-o", "package.tar.gz")
	err := mcfly_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mcfly_zip_url := "https://github.com/cantino/mcfly/archive/refs/tags/v0.9.2.zip"
	mcfly_cmd_zip := exec.Command("curl", "-L", mcfly_zip_url, "-o", "package.zip")
	err = mcfly_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mcfly_bin_url := "https://github.com/cantino/mcfly/archive/refs/tags/v0.9.2.bin"
	mcfly_cmd_bin := exec.Command("curl", "-L", mcfly_bin_url, "-o", "binary.bin")
	err = mcfly_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mcfly_src_url := "https://github.com/cantino/mcfly/archive/refs/tags/v0.9.2.src.tar.gz"
	mcfly_cmd_src := exec.Command("curl", "-L", mcfly_src_url, "-o", "source.tar.gz")
	err = mcfly_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mcfly_cmd_direct := exec.Command("./binary")
	err = mcfly_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
}
