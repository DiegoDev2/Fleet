package main

// LlvmAT17 - Next-gen compiler infrastructure
// Homepage: https://llvm.org/

import (
	"fmt"
	
	"os/exec"
)

func installLlvmAT17() {
	// Método 1: Descargar y extraer .tar.gz
	llvmat17_tar_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-17.0.6/llvm-project-17.0.6.src.tar.xz"
	llvmat17_cmd_tar := exec.Command("curl", "-L", llvmat17_tar_url, "-o", "package.tar.gz")
	err := llvmat17_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	llvmat17_zip_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-17.0.6/llvm-project-17.0.6.src.tar.xz"
	llvmat17_cmd_zip := exec.Command("curl", "-L", llvmat17_zip_url, "-o", "package.zip")
	err = llvmat17_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	llvmat17_bin_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-17.0.6/llvm-project-17.0.6.src.tar.xz"
	llvmat17_cmd_bin := exec.Command("curl", "-L", llvmat17_bin_url, "-o", "binary.bin")
	err = llvmat17_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	llvmat17_src_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-17.0.6/llvm-project-17.0.6.src.tar.xz"
	llvmat17_cmd_src := exec.Command("curl", "-L", llvmat17_src_url, "-o", "source.tar.gz")
	err = llvmat17_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	llvmat17_cmd_direct := exec.Command("./binary")
	err = llvmat17_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
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
