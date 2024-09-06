package main

// Detox - Utility to replace problematic characters in filenames
// Homepage: https://detox.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installDetox() {
	// Método 1: Descargar y extraer .tar.gz
	detox_tar_url := "https://github.com/dharple/detox/archive/refs/tags/v2.0.0.tar.gz"
	detox_cmd_tar := exec.Command("curl", "-L", detox_tar_url, "-o", "package.tar.gz")
	err := detox_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	detox_zip_url := "https://github.com/dharple/detox/archive/refs/tags/v2.0.0.zip"
	detox_cmd_zip := exec.Command("curl", "-L", detox_zip_url, "-o", "package.zip")
	err = detox_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	detox_bin_url := "https://github.com/dharple/detox/archive/refs/tags/v2.0.0.bin"
	detox_cmd_bin := exec.Command("curl", "-L", detox_bin_url, "-o", "binary.bin")
	err = detox_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	detox_src_url := "https://github.com/dharple/detox/archive/refs/tags/v2.0.0.src.tar.gz"
	detox_cmd_src := exec.Command("curl", "-L", detox_src_url, "-o", "source.tar.gz")
	err = detox_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	detox_cmd_direct := exec.Command("./binary")
	err = detox_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
}
