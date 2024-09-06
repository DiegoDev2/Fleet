package main

// Libswiftnav - C library implementing GNSS related functions and algorithms
// Homepage: https://github.com/swift-nav/libswiftnav

import (
	"fmt"
	
	"os/exec"
)

func installLibswiftnav() {
	// Método 1: Descargar y extraer .tar.gz
	libswiftnav_tar_url := "https://github.com/swift-nav/libswiftnav/archive/refs/tags/v2.4.2.tar.gz"
	libswiftnav_cmd_tar := exec.Command("curl", "-L", libswiftnav_tar_url, "-o", "package.tar.gz")
	err := libswiftnav_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libswiftnav_zip_url := "https://github.com/swift-nav/libswiftnav/archive/refs/tags/v2.4.2.zip"
	libswiftnav_cmd_zip := exec.Command("curl", "-L", libswiftnav_zip_url, "-o", "package.zip")
	err = libswiftnav_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libswiftnav_bin_url := "https://github.com/swift-nav/libswiftnav/archive/refs/tags/v2.4.2.bin"
	libswiftnav_cmd_bin := exec.Command("curl", "-L", libswiftnav_bin_url, "-o", "binary.bin")
	err = libswiftnav_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libswiftnav_src_url := "https://github.com/swift-nav/libswiftnav/archive/refs/tags/v2.4.2.src.tar.gz"
	libswiftnav_cmd_src := exec.Command("curl", "-L", libswiftnav_src_url, "-o", "source.tar.gz")
	err = libswiftnav_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libswiftnav_cmd_direct := exec.Command("./binary")
	err = libswiftnav_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
