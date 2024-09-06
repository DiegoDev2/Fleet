package main

// Exodriver - Thin interface to LabJack devices
// Homepage: https://labjack.com/support/linux-and-mac-os-x-drivers

import (
	"fmt"
	
	"os/exec"
)

func installExodriver() {
	// Método 1: Descargar y extraer .tar.gz
	exodriver_tar_url := "https://github.com/labjack/exodriver/archive/refs/tags/v2.7.0.tar.gz"
	exodriver_cmd_tar := exec.Command("curl", "-L", exodriver_tar_url, "-o", "package.tar.gz")
	err := exodriver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exodriver_zip_url := "https://github.com/labjack/exodriver/archive/refs/tags/v2.7.0.zip"
	exodriver_cmd_zip := exec.Command("curl", "-L", exodriver_zip_url, "-o", "package.zip")
	err = exodriver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exodriver_bin_url := "https://github.com/labjack/exodriver/archive/refs/tags/v2.7.0.bin"
	exodriver_cmd_bin := exec.Command("curl", "-L", exodriver_bin_url, "-o", "binary.bin")
	err = exodriver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exodriver_src_url := "https://github.com/labjack/exodriver/archive/refs/tags/v2.7.0.src.tar.gz"
	exodriver_cmd_src := exec.Command("curl", "-L", exodriver_src_url, "-o", "source.tar.gz")
	err = exodriver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exodriver_cmd_direct := exec.Command("./binary")
	err = exodriver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
