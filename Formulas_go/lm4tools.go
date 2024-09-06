package main

// Lm4tools - Tools for TI Stellaris Launchpad boards
// Homepage: https://github.com/utzig/lm4tools

import (
	"fmt"
	
	"os/exec"
)

func installLm4tools() {
	// Método 1: Descargar y extraer .tar.gz
	lm4tools_tar_url := "https://github.com/utzig/lm4tools/archive/refs/tags/v0.1.3.tar.gz"
	lm4tools_cmd_tar := exec.Command("curl", "-L", lm4tools_tar_url, "-o", "package.tar.gz")
	err := lm4tools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lm4tools_zip_url := "https://github.com/utzig/lm4tools/archive/refs/tags/v0.1.3.zip"
	lm4tools_cmd_zip := exec.Command("curl", "-L", lm4tools_zip_url, "-o", "package.zip")
	err = lm4tools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lm4tools_bin_url := "https://github.com/utzig/lm4tools/archive/refs/tags/v0.1.3.bin"
	lm4tools_cmd_bin := exec.Command("curl", "-L", lm4tools_bin_url, "-o", "binary.bin")
	err = lm4tools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lm4tools_src_url := "https://github.com/utzig/lm4tools/archive/refs/tags/v0.1.3.src.tar.gz"
	lm4tools_cmd_src := exec.Command("curl", "-L", lm4tools_src_url, "-o", "source.tar.gz")
	err = lm4tools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lm4tools_cmd_direct := exec.Command("./binary")
	err = lm4tools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
}
