package main

// Nodeenv - Node.js virtual environment builder
// Homepage: https://github.com/ekalinin/nodeenv

import (
	"fmt"
	
	"os/exec"
)

func installNodeenv() {
	// Método 1: Descargar y extraer .tar.gz
	nodeenv_tar_url := "https://github.com/ekalinin/nodeenv/archive/refs/tags/1.9.1.tar.gz"
	nodeenv_cmd_tar := exec.Command("curl", "-L", nodeenv_tar_url, "-o", "package.tar.gz")
	err := nodeenv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nodeenv_zip_url := "https://github.com/ekalinin/nodeenv/archive/refs/tags/1.9.1.zip"
	nodeenv_cmd_zip := exec.Command("curl", "-L", nodeenv_zip_url, "-o", "package.zip")
	err = nodeenv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nodeenv_bin_url := "https://github.com/ekalinin/nodeenv/archive/refs/tags/1.9.1.bin"
	nodeenv_cmd_bin := exec.Command("curl", "-L", nodeenv_bin_url, "-o", "binary.bin")
	err = nodeenv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nodeenv_src_url := "https://github.com/ekalinin/nodeenv/archive/refs/tags/1.9.1.src.tar.gz"
	nodeenv_cmd_src := exec.Command("curl", "-L", nodeenv_src_url, "-o", "source.tar.gz")
	err = nodeenv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nodeenv_cmd_direct := exec.Command("./binary")
	err = nodeenv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
