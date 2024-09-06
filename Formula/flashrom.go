package main

// Flashrom - Identify, read, write, verify, and erase flash chips
// Homepage: https://flashrom.org/

import (
	"fmt"
	
	"os/exec"
)

func installFlashrom() {
	// Método 1: Descargar y extraer .tar.gz
	flashrom_tar_url := "https://download.flashrom.org/releases/flashrom-1.4.0.tar.xz"
	flashrom_cmd_tar := exec.Command("curl", "-L", flashrom_tar_url, "-o", "package.tar.gz")
	err := flashrom_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flashrom_zip_url := "https://download.flashrom.org/releases/flashrom-1.4.0.tar.xz"
	flashrom_cmd_zip := exec.Command("curl", "-L", flashrom_zip_url, "-o", "package.zip")
	err = flashrom_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flashrom_bin_url := "https://download.flashrom.org/releases/flashrom-1.4.0.tar.xz"
	flashrom_cmd_bin := exec.Command("curl", "-L", flashrom_bin_url, "-o", "binary.bin")
	err = flashrom_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flashrom_src_url := "https://download.flashrom.org/releases/flashrom-1.4.0.tar.xz"
	flashrom_cmd_src := exec.Command("curl", "-L", flashrom_src_url, "-o", "source.tar.gz")
	err = flashrom_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flashrom_cmd_direct := exec.Command("./binary")
	err = flashrom_cmd_direct.Run()
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
}
