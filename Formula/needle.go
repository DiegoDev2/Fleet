package main

// Needle - Compile-time safe Swift dependency injection framework with real code
// Homepage: https://github.com/uber/needle

import (
	"fmt"
	
	"os/exec"
)

func installNeedle() {
	// Método 1: Descargar y extraer .tar.gz
	needle_tar_url := "https://github.com/uber/needle/archive/refs/tags/v0.24.0.tar.gz"
	needle_cmd_tar := exec.Command("curl", "-L", needle_tar_url, "-o", "package.tar.gz")
	err := needle_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	needle_zip_url := "https://github.com/uber/needle/archive/refs/tags/v0.24.0.zip"
	needle_cmd_zip := exec.Command("curl", "-L", needle_zip_url, "-o", "package.zip")
	err = needle_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	needle_bin_url := "https://github.com/uber/needle/archive/refs/tags/v0.24.0.bin"
	needle_cmd_bin := exec.Command("curl", "-L", needle_bin_url, "-o", "binary.bin")
	err = needle_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	needle_src_url := "https://github.com/uber/needle/archive/refs/tags/v0.24.0.src.tar.gz"
	needle_cmd_src := exec.Command("curl", "-L", needle_src_url, "-o", "source.tar.gz")
	err = needle_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	needle_cmd_direct := exec.Command("./binary")
	err = needle_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
