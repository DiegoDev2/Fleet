package main

// Mspdebug - Debugger for use with MSP430 MCUs
// Homepage: https://dlbeer.co.nz/mspdebug/

import (
	"fmt"
	
	"os/exec"
)

func installMspdebug() {
	// Método 1: Descargar y extraer .tar.gz
	mspdebug_tar_url := "https://github.com/dlbeer/mspdebug/archive/refs/tags/v0.25.tar.gz"
	mspdebug_cmd_tar := exec.Command("curl", "-L", mspdebug_tar_url, "-o", "package.tar.gz")
	err := mspdebug_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mspdebug_zip_url := "https://github.com/dlbeer/mspdebug/archive/refs/tags/v0.25.zip"
	mspdebug_cmd_zip := exec.Command("curl", "-L", mspdebug_zip_url, "-o", "package.zip")
	err = mspdebug_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mspdebug_bin_url := "https://github.com/dlbeer/mspdebug/archive/refs/tags/v0.25.bin"
	mspdebug_cmd_bin := exec.Command("curl", "-L", mspdebug_bin_url, "-o", "binary.bin")
	err = mspdebug_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mspdebug_src_url := "https://github.com/dlbeer/mspdebug/archive/refs/tags/v0.25.src.tar.gz"
	mspdebug_cmd_src := exec.Command("curl", "-L", mspdebug_src_url, "-o", "source.tar.gz")
	err = mspdebug_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mspdebug_cmd_direct := exec.Command("./binary")
	err = mspdebug_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: hidapi")
	exec.Command("latte", "install", "hidapi").Run()
	fmt.Println("Instalando dependencia: libusb-compat")
	exec.Command("latte", "install", "libusb-compat").Run()
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
