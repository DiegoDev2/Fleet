package main

// Libcuefile - Library to work with CUE files
// Homepage: https://www.musepack.net/

import (
	"fmt"
	
	"os/exec"
)

func installLibcuefile() {
	// Método 1: Descargar y extraer .tar.gz
	libcuefile_tar_url := "https://files.musepack.net/source/libcuefile_r475.tar.gz"
	libcuefile_cmd_tar := exec.Command("curl", "-L", libcuefile_tar_url, "-o", "package.tar.gz")
	err := libcuefile_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcuefile_zip_url := "https://files.musepack.net/source/libcuefile_r475.zip"
	libcuefile_cmd_zip := exec.Command("curl", "-L", libcuefile_zip_url, "-o", "package.zip")
	err = libcuefile_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcuefile_bin_url := "https://files.musepack.net/source/libcuefile_r475.bin"
	libcuefile_cmd_bin := exec.Command("curl", "-L", libcuefile_bin_url, "-o", "binary.bin")
	err = libcuefile_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcuefile_src_url := "https://files.musepack.net/source/libcuefile_r475.src.tar.gz"
	libcuefile_cmd_src := exec.Command("curl", "-L", libcuefile_src_url, "-o", "source.tar.gz")
	err = libcuefile_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcuefile_cmd_direct := exec.Command("./binary")
	err = libcuefile_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
