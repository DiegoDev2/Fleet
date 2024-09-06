package main

// Libversion - Advanced version string comparison library
// Homepage: https://github.com/repology/libversion

import (
	"fmt"
	
	"os/exec"
)

func installLibversion() {
	// Método 1: Descargar y extraer .tar.gz
	libversion_tar_url := "https://github.com/repology/libversion/archive/refs/tags/3.0.3.tar.gz"
	libversion_cmd_tar := exec.Command("curl", "-L", libversion_tar_url, "-o", "package.tar.gz")
	err := libversion_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libversion_zip_url := "https://github.com/repology/libversion/archive/refs/tags/3.0.3.zip"
	libversion_cmd_zip := exec.Command("curl", "-L", libversion_zip_url, "-o", "package.zip")
	err = libversion_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libversion_bin_url := "https://github.com/repology/libversion/archive/refs/tags/3.0.3.bin"
	libversion_cmd_bin := exec.Command("curl", "-L", libversion_bin_url, "-o", "binary.bin")
	err = libversion_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libversion_src_url := "https://github.com/repology/libversion/archive/refs/tags/3.0.3.src.tar.gz"
	libversion_cmd_src := exec.Command("curl", "-L", libversion_src_url, "-o", "source.tar.gz")
	err = libversion_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libversion_cmd_direct := exec.Command("./binary")
	err = libversion_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
