package main

// LlvmAT16 - Next-gen compiler infrastructure
// Homepage: https://llvm.org/

import (
	"fmt"
	
	"os/exec"
)

func installLlvmAT16() {
	// Método 1: Descargar y extraer .tar.gz
	llvmat16_tar_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-16.0.6/llvm-project-16.0.6.src.tar.xz"
	llvmat16_cmd_tar := exec.Command("curl", "-L", llvmat16_tar_url, "-o", "package.tar.gz")
	err := llvmat16_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	llvmat16_zip_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-16.0.6/llvm-project-16.0.6.src.tar.xz"
	llvmat16_cmd_zip := exec.Command("curl", "-L", llvmat16_zip_url, "-o", "package.zip")
	err = llvmat16_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	llvmat16_bin_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-16.0.6/llvm-project-16.0.6.src.tar.xz"
	llvmat16_cmd_bin := exec.Command("curl", "-L", llvmat16_bin_url, "-o", "binary.bin")
	err = llvmat16_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	llvmat16_src_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-16.0.6/llvm-project-16.0.6.src.tar.xz"
	llvmat16_cmd_src := exec.Command("curl", "-L", llvmat16_src_url, "-o", "source.tar.gz")
	err = llvmat16_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	llvmat16_cmd_direct := exec.Command("./binary")
	err = llvmat16_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: binutils")
exec.Command("latte", "install", "binutils")
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
}
