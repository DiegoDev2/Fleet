package main

// Libdivide - Optimized integer division
// Homepage: https://libdivide.com

import (
	"fmt"
	
	"os/exec"
)

func installLibdivide() {
	// Método 1: Descargar y extraer .tar.gz
	libdivide_tar_url := "https://github.com/ridiculousfish/libdivide/archive/refs/tags/v5.1.tar.gz"
	libdivide_cmd_tar := exec.Command("curl", "-L", libdivide_tar_url, "-o", "package.tar.gz")
	err := libdivide_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdivide_zip_url := "https://github.com/ridiculousfish/libdivide/archive/refs/tags/v5.1.zip"
	libdivide_cmd_zip := exec.Command("curl", "-L", libdivide_zip_url, "-o", "package.zip")
	err = libdivide_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdivide_bin_url := "https://github.com/ridiculousfish/libdivide/archive/refs/tags/v5.1.bin"
	libdivide_cmd_bin := exec.Command("curl", "-L", libdivide_bin_url, "-o", "binary.bin")
	err = libdivide_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdivide_src_url := "https://github.com/ridiculousfish/libdivide/archive/refs/tags/v5.1.src.tar.gz"
	libdivide_cmd_src := exec.Command("curl", "-L", libdivide_src_url, "-o", "source.tar.gz")
	err = libdivide_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdivide_cmd_direct := exec.Command("./binary")
	err = libdivide_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
