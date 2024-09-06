package main

// Libsoundio - Cross-platform audio input and output
// Homepage: http://libsound.io

import (
	"fmt"
	
	"os/exec"
)

func installLibsoundio() {
	// Método 1: Descargar y extraer .tar.gz
	libsoundio_tar_url := "https://github.com/andrewrk/libsoundio/archive/refs/tags/2.0.1-7.tar.gz"
	libsoundio_cmd_tar := exec.Command("curl", "-L", libsoundio_tar_url, "-o", "package.tar.gz")
	err := libsoundio_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsoundio_zip_url := "https://github.com/andrewrk/libsoundio/archive/refs/tags/2.0.1-7.zip"
	libsoundio_cmd_zip := exec.Command("curl", "-L", libsoundio_zip_url, "-o", "package.zip")
	err = libsoundio_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsoundio_bin_url := "https://github.com/andrewrk/libsoundio/archive/refs/tags/2.0.1-7.bin"
	libsoundio_cmd_bin := exec.Command("curl", "-L", libsoundio_bin_url, "-o", "binary.bin")
	err = libsoundio_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsoundio_src_url := "https://github.com/andrewrk/libsoundio/archive/refs/tags/2.0.1-7.src.tar.gz"
	libsoundio_cmd_src := exec.Command("curl", "-L", libsoundio_src_url, "-o", "source.tar.gz")
	err = libsoundio_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsoundio_cmd_direct := exec.Command("./binary")
	err = libsoundio_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
