package main

// Libclc - Implementation of the library requirements of the OpenCL C programming language
// Homepage: https://libclc.llvm.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibclc() {
	// Método 1: Descargar y extraer .tar.gz
	libclc_tar_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/libclc-18.1.8.src.tar.xz"
	libclc_cmd_tar := exec.Command("curl", "-L", libclc_tar_url, "-o", "package.tar.gz")
	err := libclc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libclc_zip_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/libclc-18.1.8.src.tar.xz"
	libclc_cmd_zip := exec.Command("curl", "-L", libclc_zip_url, "-o", "package.zip")
	err = libclc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libclc_bin_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/libclc-18.1.8.src.tar.xz"
	libclc_cmd_bin := exec.Command("curl", "-L", libclc_bin_url, "-o", "binary.bin")
	err = libclc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libclc_src_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/libclc-18.1.8.src.tar.xz"
	libclc_cmd_src := exec.Command("curl", "-L", libclc_src_url, "-o", "source.tar.gz")
	err = libclc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libclc_cmd_direct := exec.Command("./binary")
	err = libclc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: spirv-llvm-translator")
exec.Command("latte", "install", "spirv-llvm-translator")
}
