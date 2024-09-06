package main

// Libvisual - Audio Visualization tool and library
// Homepage: https://github.com/Libvisual/libvisual

import (
	"fmt"
	
	"os/exec"
)

func installLibvisual() {
	// Método 1: Descargar y extraer .tar.gz
	libvisual_tar_url := "https://github.com/Libvisual/libvisual/releases/download/libvisual-0.4.2/libvisual-0.4.2.tar.gz"
	libvisual_cmd_tar := exec.Command("curl", "-L", libvisual_tar_url, "-o", "package.tar.gz")
	err := libvisual_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libvisual_zip_url := "https://github.com/Libvisual/libvisual/releases/download/libvisual-0.4.2/libvisual-0.4.2.zip"
	libvisual_cmd_zip := exec.Command("curl", "-L", libvisual_zip_url, "-o", "package.zip")
	err = libvisual_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libvisual_bin_url := "https://github.com/Libvisual/libvisual/releases/download/libvisual-0.4.2/libvisual-0.4.2.bin"
	libvisual_cmd_bin := exec.Command("curl", "-L", libvisual_bin_url, "-o", "binary.bin")
	err = libvisual_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libvisual_src_url := "https://github.com/Libvisual/libvisual/releases/download/libvisual-0.4.2/libvisual-0.4.2.src.tar.gz"
	libvisual_cmd_src := exec.Command("curl", "-L", libvisual_src_url, "-o", "source.tar.gz")
	err = libvisual_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libvisual_cmd_direct := exec.Command("./binary")
	err = libvisual_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: sdl12-compat")
	exec.Command("latte", "install", "sdl12-compat").Run()
}
