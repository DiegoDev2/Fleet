package main

// Qdmr - Codeplug programming tool for DMR radios
// Homepage: https://dm3mat.darc.de/qdmr/

import (
	"fmt"
	
	"os/exec"
)

func installQdmr() {
	// Método 1: Descargar y extraer .tar.gz
	qdmr_tar_url := "https://github.com/hmatuschek/qdmr/archive/refs/tags/v0.12.0.tar.gz"
	qdmr_cmd_tar := exec.Command("curl", "-L", qdmr_tar_url, "-o", "package.tar.gz")
	err := qdmr_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qdmr_zip_url := "https://github.com/hmatuschek/qdmr/archive/refs/tags/v0.12.0.zip"
	qdmr_cmd_zip := exec.Command("curl", "-L", qdmr_zip_url, "-o", "package.zip")
	err = qdmr_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qdmr_bin_url := "https://github.com/hmatuschek/qdmr/archive/refs/tags/v0.12.0.bin"
	qdmr_cmd_bin := exec.Command("curl", "-L", qdmr_bin_url, "-o", "binary.bin")
	err = qdmr_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qdmr_src_url := "https://github.com/hmatuschek/qdmr/archive/refs/tags/v0.12.0.src.tar.gz"
	qdmr_cmd_src := exec.Command("curl", "-L", qdmr_src_url, "-o", "source.tar.gz")
	err = qdmr_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qdmr_cmd_direct := exec.Command("./binary")
	err = qdmr_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libusb")
exec.Command("latte", "install", "libusb")
	fmt.Println("Instalando dependencia: qt@5")
exec.Command("latte", "install", "qt@5")
	fmt.Println("Instalando dependencia: yaml-cpp")
exec.Command("latte", "install", "yaml-cpp")
}
