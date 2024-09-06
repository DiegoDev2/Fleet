package main

// Libunwind - C API for determining the call-chain of a program
// Homepage: https://www.nongnu.org/libunwind/

import (
	"fmt"
	
	"os/exec"
)

func installLibunwind() {
	// Método 1: Descargar y extraer .tar.gz
	libunwind_tar_url := "https://github.com/libunwind/libunwind/releases/download/v1.8.1/libunwind-1.8.1.tar.gz"
	libunwind_cmd_tar := exec.Command("curl", "-L", libunwind_tar_url, "-o", "package.tar.gz")
	err := libunwind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libunwind_zip_url := "https://github.com/libunwind/libunwind/releases/download/v1.8.1/libunwind-1.8.1.zip"
	libunwind_cmd_zip := exec.Command("curl", "-L", libunwind_zip_url, "-o", "package.zip")
	err = libunwind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libunwind_bin_url := "https://github.com/libunwind/libunwind/releases/download/v1.8.1/libunwind-1.8.1.bin"
	libunwind_cmd_bin := exec.Command("curl", "-L", libunwind_bin_url, "-o", "binary.bin")
	err = libunwind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libunwind_src_url := "https://github.com/libunwind/libunwind/releases/download/v1.8.1/libunwind-1.8.1.src.tar.gz"
	libunwind_cmd_src := exec.Command("curl", "-L", libunwind_src_url, "-o", "source.tar.gz")
	err = libunwind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libunwind_cmd_direct := exec.Command("./binary")
	err = libunwind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: zlib")
	exec.Command("latte", "install", "zlib").Run()
}
