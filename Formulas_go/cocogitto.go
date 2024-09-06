package main

// Cocogitto - Conventional Commits toolbox
// Homepage: https://docs.cocogitto.io/

import (
	"fmt"
	
	"os/exec"
)

func installCocogitto() {
	// Método 1: Descargar y extraer .tar.gz
	cocogitto_tar_url := "https://github.com/cocogitto/cocogitto/archive/refs/tags/6.1.0.tar.gz"
	cocogitto_cmd_tar := exec.Command("curl", "-L", cocogitto_tar_url, "-o", "package.tar.gz")
	err := cocogitto_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cocogitto_zip_url := "https://github.com/cocogitto/cocogitto/archive/refs/tags/6.1.0.zip"
	cocogitto_cmd_zip := exec.Command("curl", "-L", cocogitto_zip_url, "-o", "package.zip")
	err = cocogitto_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cocogitto_bin_url := "https://github.com/cocogitto/cocogitto/archive/refs/tags/6.1.0.bin"
	cocogitto_cmd_bin := exec.Command("curl", "-L", cocogitto_bin_url, "-o", "binary.bin")
	err = cocogitto_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cocogitto_src_url := "https://github.com/cocogitto/cocogitto/archive/refs/tags/6.1.0.src.tar.gz"
	cocogitto_cmd_src := exec.Command("curl", "-L", cocogitto_src_url, "-o", "source.tar.gz")
	err = cocogitto_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cocogitto_cmd_direct := exec.Command("./binary")
	err = cocogitto_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: libgit2@1.7")
exec.Command("latte", "install", "libgit2@1.7")
}
