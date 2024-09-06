package main

// Libcbor - CBOR protocol implementation for C and others
// Homepage: https://github.com/PJK/libcbor

import (
	"fmt"
	
	"os/exec"
)

func installLibcbor() {
	// Método 1: Descargar y extraer .tar.gz
	libcbor_tar_url := "https://github.com/PJK/libcbor/archive/refs/tags/v0.11.0.tar.gz"
	libcbor_cmd_tar := exec.Command("curl", "-L", libcbor_tar_url, "-o", "package.tar.gz")
	err := libcbor_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libcbor_zip_url := "https://github.com/PJK/libcbor/archive/refs/tags/v0.11.0.zip"
	libcbor_cmd_zip := exec.Command("curl", "-L", libcbor_zip_url, "-o", "package.zip")
	err = libcbor_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libcbor_bin_url := "https://github.com/PJK/libcbor/archive/refs/tags/v0.11.0.bin"
	libcbor_cmd_bin := exec.Command("curl", "-L", libcbor_bin_url, "-o", "binary.bin")
	err = libcbor_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libcbor_src_url := "https://github.com/PJK/libcbor/archive/refs/tags/v0.11.0.src.tar.gz"
	libcbor_cmd_src := exec.Command("curl", "-L", libcbor_src_url, "-o", "source.tar.gz")
	err = libcbor_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libcbor_cmd_direct := exec.Command("./binary")
	err = libcbor_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
