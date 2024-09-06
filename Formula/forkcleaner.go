package main

// ForkCleaner - Cleans up old and inactive forks on your GitHub account
// Homepage: https://github.com/caarlos0/fork-cleaner

import (
	"fmt"
	
	"os/exec"
)

func installForkCleaner() {
	// Método 1: Descargar y extraer .tar.gz
	forkcleaner_tar_url := "https://github.com/caarlos0/fork-cleaner/archive/refs/tags/v2.3.1.tar.gz"
	forkcleaner_cmd_tar := exec.Command("curl", "-L", forkcleaner_tar_url, "-o", "package.tar.gz")
	err := forkcleaner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	forkcleaner_zip_url := "https://github.com/caarlos0/fork-cleaner/archive/refs/tags/v2.3.1.zip"
	forkcleaner_cmd_zip := exec.Command("curl", "-L", forkcleaner_zip_url, "-o", "package.zip")
	err = forkcleaner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	forkcleaner_bin_url := "https://github.com/caarlos0/fork-cleaner/archive/refs/tags/v2.3.1.bin"
	forkcleaner_cmd_bin := exec.Command("curl", "-L", forkcleaner_bin_url, "-o", "binary.bin")
	err = forkcleaner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	forkcleaner_src_url := "https://github.com/caarlos0/fork-cleaner/archive/refs/tags/v2.3.1.src.tar.gz"
	forkcleaner_cmd_src := exec.Command("curl", "-L", forkcleaner_src_url, "-o", "source.tar.gz")
	err = forkcleaner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	forkcleaner_cmd_direct := exec.Command("./binary")
	err = forkcleaner_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
