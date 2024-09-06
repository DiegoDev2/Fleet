package main

// Fw - Workspace productivity booster
// Homepage: https://github.com/brocode/fw

import (
	"fmt"
	
	"os/exec"
)

func installFw() {
	// Método 1: Descargar y extraer .tar.gz
	fw_tar_url := "https://github.com/brocode/fw/archive/refs/tags/v2.19.1.tar.gz"
	fw_cmd_tar := exec.Command("curl", "-L", fw_tar_url, "-o", "package.tar.gz")
	err := fw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fw_zip_url := "https://github.com/brocode/fw/archive/refs/tags/v2.19.1.zip"
	fw_cmd_zip := exec.Command("curl", "-L", fw_zip_url, "-o", "package.zip")
	err = fw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fw_bin_url := "https://github.com/brocode/fw/archive/refs/tags/v2.19.1.bin"
	fw_cmd_bin := exec.Command("curl", "-L", fw_bin_url, "-o", "binary.bin")
	err = fw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fw_src_url := "https://github.com/brocode/fw/archive/refs/tags/v2.19.1.src.tar.gz"
	fw_cmd_src := exec.Command("curl", "-L", fw_src_url, "-o", "source.tar.gz")
	err = fw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fw_cmd_direct := exec.Command("./binary")
	err = fw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
