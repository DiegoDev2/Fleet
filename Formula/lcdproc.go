package main

// Lcdproc - Display real-time system information on a LCD
// Homepage: https://www.lcdproc.org/

import (
	"fmt"
	
	"os/exec"
)

func installLcdproc() {
	// Método 1: Descargar y extraer .tar.gz
	lcdproc_tar_url := "https://github.com/lcdproc/lcdproc/releases/download/v0.5.9/lcdproc-0.5.9.tar.gz"
	lcdproc_cmd_tar := exec.Command("curl", "-L", lcdproc_tar_url, "-o", "package.tar.gz")
	err := lcdproc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lcdproc_zip_url := "https://github.com/lcdproc/lcdproc/releases/download/v0.5.9/lcdproc-0.5.9.zip"
	lcdproc_cmd_zip := exec.Command("curl", "-L", lcdproc_zip_url, "-o", "package.zip")
	err = lcdproc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lcdproc_bin_url := "https://github.com/lcdproc/lcdproc/releases/download/v0.5.9/lcdproc-0.5.9.bin"
	lcdproc_cmd_bin := exec.Command("curl", "-L", lcdproc_bin_url, "-o", "binary.bin")
	err = lcdproc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lcdproc_src_url := "https://github.com/lcdproc/lcdproc/releases/download/v0.5.9/lcdproc-0.5.9.src.tar.gz"
	lcdproc_cmd_src := exec.Command("curl", "-L", lcdproc_src_url, "-o", "source.tar.gz")
	err = lcdproc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lcdproc_cmd_direct := exec.Command("./binary")
	err = lcdproc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libftdi")
	exec.Command("latte", "install", "libftdi").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
	fmt.Println("Instalando dependencia: libusb-compat")
	exec.Command("latte", "install", "libusb-compat").Run()
}
