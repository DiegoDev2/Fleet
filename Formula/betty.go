package main

// Betty - English-like interface for the command-line
// Homepage: https://github.com/pickhardt/betty

import (
	"fmt"
	
	"os/exec"
)

func installBetty() {
	// Método 1: Descargar y extraer .tar.gz
	betty_tar_url := "https://github.com/pickhardt/betty/archive/refs/tags/v0.1.7.tar.gz"
	betty_cmd_tar := exec.Command("curl", "-L", betty_tar_url, "-o", "package.tar.gz")
	err := betty_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	betty_zip_url := "https://github.com/pickhardt/betty/archive/refs/tags/v0.1.7.zip"
	betty_cmd_zip := exec.Command("curl", "-L", betty_zip_url, "-o", "package.zip")
	err = betty_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	betty_bin_url := "https://github.com/pickhardt/betty/archive/refs/tags/v0.1.7.bin"
	betty_cmd_bin := exec.Command("curl", "-L", betty_bin_url, "-o", "binary.bin")
	err = betty_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	betty_src_url := "https://github.com/pickhardt/betty/archive/refs/tags/v0.1.7.src.tar.gz"
	betty_cmd_src := exec.Command("curl", "-L", betty_src_url, "-o", "source.tar.gz")
	err = betty_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	betty_cmd_direct := exec.Command("./binary")
	err = betty_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
