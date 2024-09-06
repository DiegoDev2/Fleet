package main

// Stlink - STM32 discovery line Linux programmer
// Homepage: https://github.com/stlink-org/stlink

import (
	"fmt"
	
	"os/exec"
)

func installStlink() {
	// Método 1: Descargar y extraer .tar.gz
	stlink_tar_url := "https://github.com/stlink-org/stlink/archive/refs/tags/v1.8.0.tar.gz"
	stlink_cmd_tar := exec.Command("curl", "-L", stlink_tar_url, "-o", "package.tar.gz")
	err := stlink_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stlink_zip_url := "https://github.com/stlink-org/stlink/archive/refs/tags/v1.8.0.zip"
	stlink_cmd_zip := exec.Command("curl", "-L", stlink_zip_url, "-o", "package.zip")
	err = stlink_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stlink_bin_url := "https://github.com/stlink-org/stlink/archive/refs/tags/v1.8.0.bin"
	stlink_cmd_bin := exec.Command("curl", "-L", stlink_bin_url, "-o", "binary.bin")
	err = stlink_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stlink_src_url := "https://github.com/stlink-org/stlink/archive/refs/tags/v1.8.0.src.tar.gz"
	stlink_cmd_src := exec.Command("curl", "-L", stlink_src_url, "-o", "source.tar.gz")
	err = stlink_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stlink_cmd_direct := exec.Command("./binary")
	err = stlink_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}
