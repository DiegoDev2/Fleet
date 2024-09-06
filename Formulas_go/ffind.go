package main

// Ffind - Friendlier find
// Homepage: https://github.com/sjl/friendly-find

import (
	"fmt"
	
	"os/exec"
)

func installFfind() {
	// Método 1: Descargar y extraer .tar.gz
	ffind_tar_url := "https://github.com/sjl/friendly-find/archive/refs/tags/v1.0.1.tar.gz"
	ffind_cmd_tar := exec.Command("curl", "-L", ffind_tar_url, "-o", "package.tar.gz")
	err := ffind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ffind_zip_url := "https://github.com/sjl/friendly-find/archive/refs/tags/v1.0.1.zip"
	ffind_cmd_zip := exec.Command("curl", "-L", ffind_zip_url, "-o", "package.zip")
	err = ffind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ffind_bin_url := "https://github.com/sjl/friendly-find/archive/refs/tags/v1.0.1.bin"
	ffind_cmd_bin := exec.Command("curl", "-L", ffind_bin_url, "-o", "binary.bin")
	err = ffind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ffind_src_url := "https://github.com/sjl/friendly-find/archive/refs/tags/v1.0.1.src.tar.gz"
	ffind_cmd_src := exec.Command("curl", "-L", ffind_src_url, "-o", "source.tar.gz")
	err = ffind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ffind_cmd_direct := exec.Command("./binary")
	err = ffind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
