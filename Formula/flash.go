package main

// Flash - Command-line script to flash SD card images of any kind
// Homepage: https://github.com/hypriot/flash

import (
	"fmt"
	
	"os/exec"
)

func installFlash() {
	// Método 1: Descargar y extraer .tar.gz
	flash_tar_url := "https://github.com/hypriot/flash/releases/download/2.7.2/flash"
	flash_cmd_tar := exec.Command("curl", "-L", flash_tar_url, "-o", "package.tar.gz")
	err := flash_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	flash_zip_url := "https://github.com/hypriot/flash/releases/download/2.7.2/flash"
	flash_cmd_zip := exec.Command("curl", "-L", flash_zip_url, "-o", "package.zip")
	err = flash_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	flash_bin_url := "https://github.com/hypriot/flash/releases/download/2.7.2/flash"
	flash_cmd_bin := exec.Command("curl", "-L", flash_bin_url, "-o", "binary.bin")
	err = flash_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	flash_src_url := "https://github.com/hypriot/flash/releases/download/2.7.2/flash"
	flash_cmd_src := exec.Command("curl", "-L", flash_src_url, "-o", "source.tar.gz")
	err = flash_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	flash_cmd_direct := exec.Command("./binary")
	err = flash_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
