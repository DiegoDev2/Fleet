package main

// Smlpkg - Package manager for Standard ML libraries and programs
// Homepage: https://github.com/diku-dk/smlpkg

import (
	"fmt"
	
	"os/exec"
)

func installSmlpkg() {
	// Método 1: Descargar y extraer .tar.gz
	smlpkg_tar_url := "https://github.com/diku-dk/smlpkg/archive/refs/tags/v0.1.5.tar.gz"
	smlpkg_cmd_tar := exec.Command("curl", "-L", smlpkg_tar_url, "-o", "package.tar.gz")
	err := smlpkg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	smlpkg_zip_url := "https://github.com/diku-dk/smlpkg/archive/refs/tags/v0.1.5.zip"
	smlpkg_cmd_zip := exec.Command("curl", "-L", smlpkg_zip_url, "-o", "package.zip")
	err = smlpkg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	smlpkg_bin_url := "https://github.com/diku-dk/smlpkg/archive/refs/tags/v0.1.5.bin"
	smlpkg_cmd_bin := exec.Command("curl", "-L", smlpkg_bin_url, "-o", "binary.bin")
	err = smlpkg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	smlpkg_src_url := "https://github.com/diku-dk/smlpkg/archive/refs/tags/v0.1.5.src.tar.gz"
	smlpkg_cmd_src := exec.Command("curl", "-L", smlpkg_src_url, "-o", "source.tar.gz")
	err = smlpkg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	smlpkg_cmd_direct := exec.Command("./binary")
	err = smlpkg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mlkit")
	exec.Command("latte", "install", "mlkit").Run()
}
