package main

// Ired - Minimalistic hexadecimal editor designed to be used in scripts
// Homepage: https://github.com/radare/ired

import (
	"fmt"
	
	"os/exec"
)

func installIred() {
	// Método 1: Descargar y extraer .tar.gz
	ired_tar_url := "https://github.com/radare/ired/archive/refs/tags/0.6.tar.gz"
	ired_cmd_tar := exec.Command("curl", "-L", ired_tar_url, "-o", "package.tar.gz")
	err := ired_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ired_zip_url := "https://github.com/radare/ired/archive/refs/tags/0.6.zip"
	ired_cmd_zip := exec.Command("curl", "-L", ired_zip_url, "-o", "package.zip")
	err = ired_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ired_bin_url := "https://github.com/radare/ired/archive/refs/tags/0.6.bin"
	ired_cmd_bin := exec.Command("curl", "-L", ired_bin_url, "-o", "binary.bin")
	err = ired_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ired_src_url := "https://github.com/radare/ired/archive/refs/tags/0.6.src.tar.gz"
	ired_cmd_src := exec.Command("curl", "-L", ired_src_url, "-o", "source.tar.gz")
	err = ired_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ired_cmd_direct := exec.Command("./binary")
	err = ired_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
