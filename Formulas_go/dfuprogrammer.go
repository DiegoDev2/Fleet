package main

// DfuProgrammer - Device firmware update based USB programmer for Atmel chips
// Homepage: https://github.com/dfu-programmer/dfu-programmer

import (
	"fmt"
	
	"os/exec"
)

func installDfuProgrammer() {
	// Método 1: Descargar y extraer .tar.gz
	dfuprogrammer_tar_url := "https://github.com/dfu-programmer/dfu-programmer/releases/download/v1.1.0/dfu-programmer-1.1.0.tar.gz"
	dfuprogrammer_cmd_tar := exec.Command("curl", "-L", dfuprogrammer_tar_url, "-o", "package.tar.gz")
	err := dfuprogrammer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dfuprogrammer_zip_url := "https://github.com/dfu-programmer/dfu-programmer/releases/download/v1.1.0/dfu-programmer-1.1.0.zip"
	dfuprogrammer_cmd_zip := exec.Command("curl", "-L", dfuprogrammer_zip_url, "-o", "package.zip")
	err = dfuprogrammer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dfuprogrammer_bin_url := "https://github.com/dfu-programmer/dfu-programmer/releases/download/v1.1.0/dfu-programmer-1.1.0.bin"
	dfuprogrammer_cmd_bin := exec.Command("curl", "-L", dfuprogrammer_bin_url, "-o", "binary.bin")
	err = dfuprogrammer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dfuprogrammer_src_url := "https://github.com/dfu-programmer/dfu-programmer/releases/download/v1.1.0/dfu-programmer-1.1.0.src.tar.gz"
	dfuprogrammer_cmd_src := exec.Command("curl", "-L", dfuprogrammer_src_url, "-o", "source.tar.gz")
	err = dfuprogrammer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dfuprogrammer_cmd_direct := exec.Command("./binary")
	err = dfuprogrammer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libusb-compat")
exec.Command("latte", "install", "libusb-compat")
}
