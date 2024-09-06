package main

// Rgbds - Rednex GameBoy Development System
// Homepage: https://rgbds.gbdev.io

import (
	"fmt"
	
	"os/exec"
)

func installRgbds() {
	// Método 1: Descargar y extraer .tar.gz
	rgbds_tar_url := "https://github.com/gbdev/rgbds/archive/refs/tags/v0.8.0.tar.gz"
	rgbds_cmd_tar := exec.Command("curl", "-L", rgbds_tar_url, "-o", "package.tar.gz")
	err := rgbds_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rgbds_zip_url := "https://github.com/gbdev/rgbds/archive/refs/tags/v0.8.0.zip"
	rgbds_cmd_zip := exec.Command("curl", "-L", rgbds_zip_url, "-o", "package.zip")
	err = rgbds_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rgbds_bin_url := "https://github.com/gbdev/rgbds/archive/refs/tags/v0.8.0.bin"
	rgbds_cmd_bin := exec.Command("curl", "-L", rgbds_bin_url, "-o", "binary.bin")
	err = rgbds_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rgbds_src_url := "https://github.com/gbdev/rgbds/archive/refs/tags/v0.8.0.src.tar.gz"
	rgbds_cmd_src := exec.Command("curl", "-L", rgbds_src_url, "-o", "source.tar.gz")
	err = rgbds_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rgbds_cmd_direct := exec.Command("./binary")
	err = rgbds_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
}
