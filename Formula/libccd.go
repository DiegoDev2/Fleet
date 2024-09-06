package main

// Libccd - Collision detection between two convex shapes
// Homepage: https://github.com/danfis/libccd

import (
	"fmt"
	
	"os/exec"
)

func installLibccd() {
	// Método 1: Descargar y extraer .tar.gz
	libccd_tar_url := "https://github.com/danfis/libccd/archive/refs/tags/v2.1.tar.gz"
	libccd_cmd_tar := exec.Command("curl", "-L", libccd_tar_url, "-o", "package.tar.gz")
	err := libccd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libccd_zip_url := "https://github.com/danfis/libccd/archive/refs/tags/v2.1.zip"
	libccd_cmd_zip := exec.Command("curl", "-L", libccd_zip_url, "-o", "package.zip")
	err = libccd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libccd_bin_url := "https://github.com/danfis/libccd/archive/refs/tags/v2.1.bin"
	libccd_cmd_bin := exec.Command("curl", "-L", libccd_bin_url, "-o", "binary.bin")
	err = libccd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libccd_src_url := "https://github.com/danfis/libccd/archive/refs/tags/v2.1.src.tar.gz"
	libccd_cmd_src := exec.Command("curl", "-L", libccd_src_url, "-o", "source.tar.gz")
	err = libccd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libccd_cmd_direct := exec.Command("./binary")
	err = libccd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
