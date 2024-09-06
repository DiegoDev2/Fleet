package main

// Brook - Cross-platform strong encryption and not detectable proxy. Zero-Configuration
// Homepage: https://txthinking.github.io/brook/

import (
	"fmt"
	
	"os/exec"
)

func installBrook() {
	// Método 1: Descargar y extraer .tar.gz
	brook_tar_url := "https://github.com/txthinking/brook/archive/refs/tags/v20240606.tar.gz"
	brook_cmd_tar := exec.Command("curl", "-L", brook_tar_url, "-o", "package.tar.gz")
	err := brook_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brook_zip_url := "https://github.com/txthinking/brook/archive/refs/tags/v20240606.zip"
	brook_cmd_zip := exec.Command("curl", "-L", brook_zip_url, "-o", "package.zip")
	err = brook_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brook_bin_url := "https://github.com/txthinking/brook/archive/refs/tags/v20240606.bin"
	brook_cmd_bin := exec.Command("curl", "-L", brook_bin_url, "-o", "binary.bin")
	err = brook_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brook_src_url := "https://github.com/txthinking/brook/archive/refs/tags/v20240606.src.tar.gz"
	brook_cmd_src := exec.Command("curl", "-L", brook_src_url, "-o", "source.tar.gz")
	err = brook_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brook_cmd_direct := exec.Command("./binary")
	err = brook_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
