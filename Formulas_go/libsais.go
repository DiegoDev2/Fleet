package main

// Libsais - Fast linear time suffix array, lcp array and bwt construction
// Homepage: https://github.com/IlyaGrebnov/libsais

import (
	"fmt"
	
	"os/exec"
)

func installLibsais() {
	// Método 1: Descargar y extraer .tar.gz
	libsais_tar_url := "https://github.com/IlyaGrebnov/libsais/archive/refs/tags/v2.8.5.tar.gz"
	libsais_cmd_tar := exec.Command("curl", "-L", libsais_tar_url, "-o", "package.tar.gz")
	err := libsais_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libsais_zip_url := "https://github.com/IlyaGrebnov/libsais/archive/refs/tags/v2.8.5.zip"
	libsais_cmd_zip := exec.Command("curl", "-L", libsais_zip_url, "-o", "package.zip")
	err = libsais_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libsais_bin_url := "https://github.com/IlyaGrebnov/libsais/archive/refs/tags/v2.8.5.bin"
	libsais_cmd_bin := exec.Command("curl", "-L", libsais_bin_url, "-o", "binary.bin")
	err = libsais_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libsais_src_url := "https://github.com/IlyaGrebnov/libsais/archive/refs/tags/v2.8.5.src.tar.gz"
	libsais_cmd_src := exec.Command("curl", "-L", libsais_src_url, "-o", "source.tar.gz")
	err = libsais_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libsais_cmd_direct := exec.Command("./binary")
	err = libsais_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
