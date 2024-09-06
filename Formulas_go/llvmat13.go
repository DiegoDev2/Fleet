package main

// LlvmAT13 - Next-gen compiler infrastructure
// Homepage: https://llvm.org/

import (
	"fmt"
	
	"os/exec"
)

func installLlvmAT13() {
	// Método 1: Descargar y extraer .tar.gz
	llvmat13_tar_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-13.0.1/llvm-project-13.0.1.src.tar.xz"
	llvmat13_cmd_tar := exec.Command("curl", "-L", llvmat13_tar_url, "-o", "package.tar.gz")
	err := llvmat13_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	llvmat13_zip_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-13.0.1/llvm-project-13.0.1.src.tar.xz"
	llvmat13_cmd_zip := exec.Command("curl", "-L", llvmat13_zip_url, "-o", "package.zip")
	err = llvmat13_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	llvmat13_bin_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-13.0.1/llvm-project-13.0.1.src.tar.xz"
	llvmat13_cmd_bin := exec.Command("curl", "-L", llvmat13_bin_url, "-o", "binary.bin")
	err = llvmat13_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	llvmat13_src_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-13.0.1/llvm-project-13.0.1.src.tar.xz"
	llvmat13_cmd_src := exec.Command("curl", "-L", llvmat13_src_url, "-o", "source.tar.gz")
	err = llvmat13_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	llvmat13_cmd_direct := exec.Command("./binary")
	err = llvmat13_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: python@3.11")
exec.Command("latte", "install", "python@3.11")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: binutils")
exec.Command("latte", "install", "binutils")
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
}
