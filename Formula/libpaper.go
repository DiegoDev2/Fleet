package main

// Libpaper - Library for handling paper characteristics
// Homepage: https://github.com/rrthomas/libpaper

import (
	"fmt"
	
	"os/exec"
)

func installLibpaper() {
	// Método 1: Descargar y extraer .tar.gz
	libpaper_tar_url := "https://github.com/rrthomas/libpaper/releases/download/v2.2.5/libpaper-2.2.5.tar.gz"
	libpaper_cmd_tar := exec.Command("curl", "-L", libpaper_tar_url, "-o", "package.tar.gz")
	err := libpaper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libpaper_zip_url := "https://github.com/rrthomas/libpaper/releases/download/v2.2.5/libpaper-2.2.5.zip"
	libpaper_cmd_zip := exec.Command("curl", "-L", libpaper_zip_url, "-o", "package.zip")
	err = libpaper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libpaper_bin_url := "https://github.com/rrthomas/libpaper/releases/download/v2.2.5/libpaper-2.2.5.bin"
	libpaper_cmd_bin := exec.Command("curl", "-L", libpaper_bin_url, "-o", "binary.bin")
	err = libpaper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libpaper_src_url := "https://github.com/rrthomas/libpaper/releases/download/v2.2.5/libpaper-2.2.5.src.tar.gz"
	libpaper_cmd_src := exec.Command("curl", "-L", libpaper_src_url, "-o", "source.tar.gz")
	err = libpaper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libpaper_cmd_direct := exec.Command("./binary")
	err = libpaper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
}
