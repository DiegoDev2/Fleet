package main

// Nodebrew - Node.js version manager
// Homepage: https://github.com/hokaccha/nodebrew

import (
	"fmt"
	
	"os/exec"
)

func installNodebrew() {
	// Método 1: Descargar y extraer .tar.gz
	nodebrew_tar_url := "https://github.com/hokaccha/nodebrew/archive/refs/tags/v1.2.0.tar.gz"
	nodebrew_cmd_tar := exec.Command("curl", "-L", nodebrew_tar_url, "-o", "package.tar.gz")
	err := nodebrew_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nodebrew_zip_url := "https://github.com/hokaccha/nodebrew/archive/refs/tags/v1.2.0.zip"
	nodebrew_cmd_zip := exec.Command("curl", "-L", nodebrew_zip_url, "-o", "package.zip")
	err = nodebrew_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nodebrew_bin_url := "https://github.com/hokaccha/nodebrew/archive/refs/tags/v1.2.0.bin"
	nodebrew_cmd_bin := exec.Command("curl", "-L", nodebrew_bin_url, "-o", "binary.bin")
	err = nodebrew_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nodebrew_src_url := "https://github.com/hokaccha/nodebrew/archive/refs/tags/v1.2.0.src.tar.gz"
	nodebrew_cmd_src := exec.Command("curl", "-L", nodebrew_src_url, "-o", "source.tar.gz")
	err = nodebrew_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nodebrew_cmd_direct := exec.Command("./binary")
	err = nodebrew_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
