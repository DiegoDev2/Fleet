package main

// Mantra - Tool to hunt down API key leaks in JS files and pages
// Homepage: https://amoloht.github.io

import (
	"fmt"
	
	"os/exec"
)

func installMantra() {
	// Método 1: Descargar y extraer .tar.gz
	mantra_tar_url := "https://github.com/MrEmpy/mantra/archive/refs/tags/v2.0.tar.gz"
	mantra_cmd_tar := exec.Command("curl", "-L", mantra_tar_url, "-o", "package.tar.gz")
	err := mantra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mantra_zip_url := "https://github.com/MrEmpy/mantra/archive/refs/tags/v2.0.zip"
	mantra_cmd_zip := exec.Command("curl", "-L", mantra_zip_url, "-o", "package.zip")
	err = mantra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mantra_bin_url := "https://github.com/MrEmpy/mantra/archive/refs/tags/v2.0.bin"
	mantra_cmd_bin := exec.Command("curl", "-L", mantra_bin_url, "-o", "binary.bin")
	err = mantra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mantra_src_url := "https://github.com/MrEmpy/mantra/archive/refs/tags/v2.0.src.tar.gz"
	mantra_cmd_src := exec.Command("curl", "-L", mantra_src_url, "-o", "source.tar.gz")
	err = mantra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mantra_cmd_direct := exec.Command("./binary")
	err = mantra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
