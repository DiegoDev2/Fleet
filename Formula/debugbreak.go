package main

// Debugbreak - Break into the debugger programmatically
// Homepage: https://github.com/scottt/debugbreak

import (
	"fmt"
	
	"os/exec"
)

func installDebugbreak() {
	// Método 1: Descargar y extraer .tar.gz
	debugbreak_tar_url := "https://github.com/scottt/debugbreak/archive/refs/tags/v1.0.tar.gz"
	debugbreak_cmd_tar := exec.Command("curl", "-L", debugbreak_tar_url, "-o", "package.tar.gz")
	err := debugbreak_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	debugbreak_zip_url := "https://github.com/scottt/debugbreak/archive/refs/tags/v1.0.zip"
	debugbreak_cmd_zip := exec.Command("curl", "-L", debugbreak_zip_url, "-o", "package.zip")
	err = debugbreak_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	debugbreak_bin_url := "https://github.com/scottt/debugbreak/archive/refs/tags/v1.0.bin"
	debugbreak_cmd_bin := exec.Command("curl", "-L", debugbreak_bin_url, "-o", "binary.bin")
	err = debugbreak_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	debugbreak_src_url := "https://github.com/scottt/debugbreak/archive/refs/tags/v1.0.src.tar.gz"
	debugbreak_cmd_src := exec.Command("curl", "-L", debugbreak_src_url, "-o", "source.tar.gz")
	err = debugbreak_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	debugbreak_cmd_direct := exec.Command("./binary")
	err = debugbreak_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
