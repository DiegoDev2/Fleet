package main

// Dislocker - FUSE driver to read/write Windows' BitLocker-ed volumes
// Homepage: https://github.com/Aorimn/dislocker

import (
	"fmt"
	
	"os/exec"
)

func installDislocker() {
	// Método 1: Descargar y extraer .tar.gz
	dislocker_tar_url := "https://github.com/Aorimn/dislocker/archive/refs/tags/v0.7.3.tar.gz"
	dislocker_cmd_tar := exec.Command("curl", "-L", dislocker_tar_url, "-o", "package.tar.gz")
	err := dislocker_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dislocker_zip_url := "https://github.com/Aorimn/dislocker/archive/refs/tags/v0.7.3.zip"
	dislocker_cmd_zip := exec.Command("curl", "-L", dislocker_zip_url, "-o", "package.zip")
	err = dislocker_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dislocker_bin_url := "https://github.com/Aorimn/dislocker/archive/refs/tags/v0.7.3.bin"
	dislocker_cmd_bin := exec.Command("curl", "-L", dislocker_bin_url, "-o", "binary.bin")
	err = dislocker_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dislocker_src_url := "https://github.com/Aorimn/dislocker/archive/refs/tags/v0.7.3.src.tar.gz"
	dislocker_cmd_src := exec.Command("curl", "-L", dislocker_src_url, "-o", "source.tar.gz")
	err = dislocker_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dislocker_cmd_direct := exec.Command("./binary")
	err = dislocker_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libfuse@2")
exec.Command("latte", "install", "libfuse@2")
	fmt.Println("Instalando dependencia: mbedtls")
exec.Command("latte", "install", "mbedtls")
}
