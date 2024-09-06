package main

// Iproute2mac - CLI wrapper for basic network utilities on macOS - ip command
// Homepage: https://github.com/brona/iproute2mac

import (
	"fmt"
	
	"os/exec"
)

func installIproute2mac() {
	// Método 1: Descargar y extraer .tar.gz
	iproute2mac_tar_url := "https://github.com/brona/iproute2mac/releases/download/v1.5.4/iproute2mac-1.5.4.tar.gz"
	iproute2mac_cmd_tar := exec.Command("curl", "-L", iproute2mac_tar_url, "-o", "package.tar.gz")
	err := iproute2mac_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iproute2mac_zip_url := "https://github.com/brona/iproute2mac/releases/download/v1.5.4/iproute2mac-1.5.4.zip"
	iproute2mac_cmd_zip := exec.Command("curl", "-L", iproute2mac_zip_url, "-o", "package.zip")
	err = iproute2mac_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iproute2mac_bin_url := "https://github.com/brona/iproute2mac/releases/download/v1.5.4/iproute2mac-1.5.4.bin"
	iproute2mac_cmd_bin := exec.Command("curl", "-L", iproute2mac_bin_url, "-o", "binary.bin")
	err = iproute2mac_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iproute2mac_src_url := "https://github.com/brona/iproute2mac/releases/download/v1.5.4/iproute2mac-1.5.4.src.tar.gz"
	iproute2mac_cmd_src := exec.Command("curl", "-L", iproute2mac_src_url, "-o", "source.tar.gz")
	err = iproute2mac_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iproute2mac_cmd_direct := exec.Command("./binary")
	err = iproute2mac_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
