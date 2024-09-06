package main

// Corral - Dependency manager for the Pony language
// Homepage: https://github.com/ponylang/corral

import (
	"fmt"
	
	"os/exec"
)

func installCorral() {
	// Método 1: Descargar y extraer .tar.gz
	corral_tar_url := "https://github.com/ponylang/corral/archive/refs/tags/0.8.1.tar.gz"
	corral_cmd_tar := exec.Command("curl", "-L", corral_tar_url, "-o", "package.tar.gz")
	err := corral_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	corral_zip_url := "https://github.com/ponylang/corral/archive/refs/tags/0.8.1.zip"
	corral_cmd_zip := exec.Command("curl", "-L", corral_zip_url, "-o", "package.zip")
	err = corral_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	corral_bin_url := "https://github.com/ponylang/corral/archive/refs/tags/0.8.1.bin"
	corral_cmd_bin := exec.Command("curl", "-L", corral_bin_url, "-o", "binary.bin")
	err = corral_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	corral_src_url := "https://github.com/ponylang/corral/archive/refs/tags/0.8.1.src.tar.gz"
	corral_cmd_src := exec.Command("curl", "-L", corral_src_url, "-o", "source.tar.gz")
	err = corral_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	corral_cmd_direct := exec.Command("./binary")
	err = corral_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ponyc")
	exec.Command("latte", "install", "ponyc").Run()
}
