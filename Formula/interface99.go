package main

// Interface99 - Full-featured interfaces for C99
// Homepage: https://github.com/Hirrolot/interface99

import (
	"fmt"
	
	"os/exec"
)

func installInterface99() {
	// Método 1: Descargar y extraer .tar.gz
	interface99_tar_url := "https://github.com/Hirrolot/interface99/archive/refs/tags/v1.0.1.tar.gz"
	interface99_cmd_tar := exec.Command("curl", "-L", interface99_tar_url, "-o", "package.tar.gz")
	err := interface99_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	interface99_zip_url := "https://github.com/Hirrolot/interface99/archive/refs/tags/v1.0.1.zip"
	interface99_cmd_zip := exec.Command("curl", "-L", interface99_zip_url, "-o", "package.zip")
	err = interface99_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	interface99_bin_url := "https://github.com/Hirrolot/interface99/archive/refs/tags/v1.0.1.bin"
	interface99_cmd_bin := exec.Command("curl", "-L", interface99_bin_url, "-o", "binary.bin")
	err = interface99_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	interface99_src_url := "https://github.com/Hirrolot/interface99/archive/refs/tags/v1.0.1.src.tar.gz"
	interface99_cmd_src := exec.Command("curl", "-L", interface99_src_url, "-o", "source.tar.gz")
	err = interface99_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	interface99_cmd_direct := exec.Command("./binary")
	err = interface99_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: metalang99")
	exec.Command("latte", "install", "metalang99").Run()
}
