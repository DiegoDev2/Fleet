package main

// Has - Checks presence of various command-line tools and their versions on the path
// Homepage: https://github.com/kdabir/has

import (
	"fmt"
	
	"os/exec"
)

func installHas() {
	// Método 1: Descargar y extraer .tar.gz
	has_tar_url := "https://github.com/kdabir/has/archive/refs/tags/v1.5.0.tar.gz"
	has_cmd_tar := exec.Command("curl", "-L", has_tar_url, "-o", "package.tar.gz")
	err := has_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	has_zip_url := "https://github.com/kdabir/has/archive/refs/tags/v1.5.0.zip"
	has_cmd_zip := exec.Command("curl", "-L", has_zip_url, "-o", "package.zip")
	err = has_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	has_bin_url := "https://github.com/kdabir/has/archive/refs/tags/v1.5.0.bin"
	has_cmd_bin := exec.Command("curl", "-L", has_bin_url, "-o", "binary.bin")
	err = has_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	has_src_url := "https://github.com/kdabir/has/archive/refs/tags/v1.5.0.src.tar.gz"
	has_cmd_src := exec.Command("curl", "-L", has_src_url, "-o", "source.tar.gz")
	err = has_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	has_cmd_direct := exec.Command("./binary")
	err = has_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
