package main

// Ispc - Compiler for SIMD programming on the CPU
// Homepage: https://ispc.github.io

import (
	"fmt"
	
	"os/exec"
)

func installIspc() {
	// Método 1: Descargar y extraer .tar.gz
	ispc_tar_url := "https://github.com/ispc/ispc/archive/refs/tags/v1.24.0.tar.gz"
	ispc_cmd_tar := exec.Command("curl", "-L", ispc_tar_url, "-o", "package.tar.gz")
	err := ispc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ispc_zip_url := "https://github.com/ispc/ispc/archive/refs/tags/v1.24.0.zip"
	ispc_cmd_zip := exec.Command("curl", "-L", ispc_zip_url, "-o", "package.zip")
	err = ispc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ispc_bin_url := "https://github.com/ispc/ispc/archive/refs/tags/v1.24.0.bin"
	ispc_cmd_bin := exec.Command("curl", "-L", ispc_bin_url, "-o", "binary.bin")
	err = ispc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ispc_src_url := "https://github.com/ispc/ispc/archive/refs/tags/v1.24.0.src.tar.gz"
	ispc_cmd_src := exec.Command("curl", "-L", ispc_src_url, "-o", "source.tar.gz")
	err = ispc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ispc_cmd_direct := exec.Command("./binary")
	err = ispc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: flex")
exec.Command("latte", "install", "flex")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
}
