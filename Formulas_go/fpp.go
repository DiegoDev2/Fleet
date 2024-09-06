package main

// Fpp - CLI program that accepts piped input and presents files for selection
// Homepage: https://facebook.github.io/PathPicker/

import (
	"fmt"
	
	"os/exec"
)

func installFpp() {
	// Método 1: Descargar y extraer .tar.gz
	fpp_tar_url := "https://github.com/facebook/PathPicker/archive/refs/tags/0.9.5.tar.gz"
	fpp_cmd_tar := exec.Command("curl", "-L", fpp_tar_url, "-o", "package.tar.gz")
	err := fpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fpp_zip_url := "https://github.com/facebook/PathPicker/archive/refs/tags/0.9.5.zip"
	fpp_cmd_zip := exec.Command("curl", "-L", fpp_zip_url, "-o", "package.zip")
	err = fpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fpp_bin_url := "https://github.com/facebook/PathPicker/archive/refs/tags/0.9.5.bin"
	fpp_cmd_bin := exec.Command("curl", "-L", fpp_bin_url, "-o", "binary.bin")
	err = fpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fpp_src_url := "https://github.com/facebook/PathPicker/archive/refs/tags/0.9.5.src.tar.gz"
	fpp_cmd_src := exec.Command("curl", "-L", fpp_src_url, "-o", "source.tar.gz")
	err = fpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fpp_cmd_direct := exec.Command("./binary")
	err = fpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
