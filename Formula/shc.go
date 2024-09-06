package main

// Shc - Shell Script Compiler
// Homepage: https://neurobin.github.io/shc

import (
	"fmt"
	
	"os/exec"
)

func installShc() {
	// Método 1: Descargar y extraer .tar.gz
	shc_tar_url := "https://github.com/neurobin/shc/archive/refs/tags/4.0.3.tar.gz"
	shc_cmd_tar := exec.Command("curl", "-L", shc_tar_url, "-o", "package.tar.gz")
	err := shc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	shc_zip_url := "https://github.com/neurobin/shc/archive/refs/tags/4.0.3.zip"
	shc_cmd_zip := exec.Command("curl", "-L", shc_zip_url, "-o", "package.zip")
	err = shc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	shc_bin_url := "https://github.com/neurobin/shc/archive/refs/tags/4.0.3.bin"
	shc_cmd_bin := exec.Command("curl", "-L", shc_bin_url, "-o", "binary.bin")
	err = shc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	shc_src_url := "https://github.com/neurobin/shc/archive/refs/tags/4.0.3.src.tar.gz"
	shc_cmd_src := exec.Command("curl", "-L", shc_src_url, "-o", "source.tar.gz")
	err = shc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	shc_cmd_direct := exec.Command("./binary")
	err = shc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
