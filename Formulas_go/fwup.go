package main

// Fwup - Configurable embedded Linux firmware update creator and runner
// Homepage: https://github.com/fwup-home/fwup

import (
	"fmt"
	
	"os/exec"
)

func installFwup() {
	// Método 1: Descargar y extraer .tar.gz
	fwup_tar_url := "https://github.com/fwup-home/fwup/releases/download/v1.10.2/fwup-1.10.2.tar.gz"
	fwup_cmd_tar := exec.Command("curl", "-L", fwup_tar_url, "-o", "package.tar.gz")
	err := fwup_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fwup_zip_url := "https://github.com/fwup-home/fwup/releases/download/v1.10.2/fwup-1.10.2.zip"
	fwup_cmd_zip := exec.Command("curl", "-L", fwup_zip_url, "-o", "package.zip")
	err = fwup_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fwup_bin_url := "https://github.com/fwup-home/fwup/releases/download/v1.10.2/fwup-1.10.2.bin"
	fwup_cmd_bin := exec.Command("curl", "-L", fwup_bin_url, "-o", "binary.bin")
	err = fwup_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fwup_src_url := "https://github.com/fwup-home/fwup/releases/download/v1.10.2/fwup-1.10.2.src.tar.gz"
	fwup_cmd_src := exec.Command("curl", "-L", fwup_src_url, "-o", "source.tar.gz")
	err = fwup_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fwup_cmd_direct := exec.Command("./binary")
	err = fwup_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: confuse")
exec.Command("latte", "install", "confuse")
	fmt.Println("Instalando dependencia: libarchive")
exec.Command("latte", "install", "libarchive")
}
