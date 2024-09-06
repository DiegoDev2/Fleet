package main

// Libmonome - Library for easy interaction with monome devices
// Homepage: https://monome.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibmonome() {
	// Método 1: Descargar y extraer .tar.gz
	libmonome_tar_url := "https://github.com/monome/libmonome/archive/refs/tags/v1.4.7.tar.gz"
	libmonome_cmd_tar := exec.Command("curl", "-L", libmonome_tar_url, "-o", "package.tar.gz")
	err := libmonome_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmonome_zip_url := "https://github.com/monome/libmonome/archive/refs/tags/v1.4.7.zip"
	libmonome_cmd_zip := exec.Command("curl", "-L", libmonome_zip_url, "-o", "package.zip")
	err = libmonome_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmonome_bin_url := "https://github.com/monome/libmonome/archive/refs/tags/v1.4.7.bin"
	libmonome_cmd_bin := exec.Command("curl", "-L", libmonome_bin_url, "-o", "binary.bin")
	err = libmonome_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmonome_src_url := "https://github.com/monome/libmonome/archive/refs/tags/v1.4.7.src.tar.gz"
	libmonome_cmd_src := exec.Command("curl", "-L", libmonome_src_url, "-o", "source.tar.gz")
	err = libmonome_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmonome_cmd_direct := exec.Command("./binary")
	err = libmonome_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: liblo")
exec.Command("latte", "install", "liblo")
}
