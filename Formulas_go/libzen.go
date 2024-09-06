package main

// Libzen - Shared library for libmediainfo
// Homepage: https://github.com/MediaArea/ZenLib

import (
	"fmt"
	
	"os/exec"
)

func installLibzen() {
	// Método 1: Descargar y extraer .tar.gz
	libzen_tar_url := "https://mediaarea.net/download/source/libzen/0.4.41/libzen_0.4.41.tar.bz2"
	libzen_cmd_tar := exec.Command("curl", "-L", libzen_tar_url, "-o", "package.tar.gz")
	err := libzen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libzen_zip_url := "https://mediaarea.net/download/source/libzen/0.4.41/libzen_0.4.41.tar.bz2"
	libzen_cmd_zip := exec.Command("curl", "-L", libzen_zip_url, "-o", "package.zip")
	err = libzen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libzen_bin_url := "https://mediaarea.net/download/source/libzen/0.4.41/libzen_0.4.41.tar.bz2"
	libzen_cmd_bin := exec.Command("curl", "-L", libzen_bin_url, "-o", "binary.bin")
	err = libzen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libzen_src_url := "https://mediaarea.net/download/source/libzen/0.4.41/libzen_0.4.41.tar.bz2"
	libzen_cmd_src := exec.Command("curl", "-L", libzen_src_url, "-o", "source.tar.gz")
	err = libzen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libzen_cmd_direct := exec.Command("./binary")
	err = libzen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
