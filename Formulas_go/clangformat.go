package main

// ClangFormat - Formatting tools for C, C++, Obj-C, Java, JavaScript, TypeScript
// Homepage: https://clang.llvm.org/docs/ClangFormat.html

import (
	"fmt"
	
	"os/exec"
)

func installClangFormat() {
	// Método 1: Descargar y extraer .tar.gz
	clangformat_tar_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/llvm-18.1.8.src.tar.xz"
	clangformat_cmd_tar := exec.Command("curl", "-L", clangformat_tar_url, "-o", "package.tar.gz")
	err := clangformat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clangformat_zip_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/llvm-18.1.8.src.tar.xz"
	clangformat_cmd_zip := exec.Command("curl", "-L", clangformat_zip_url, "-o", "package.zip")
	err = clangformat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clangformat_bin_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/llvm-18.1.8.src.tar.xz"
	clangformat_cmd_bin := exec.Command("curl", "-L", clangformat_bin_url, "-o", "binary.bin")
	err = clangformat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clangformat_src_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-18.1.8/llvm-18.1.8.src.tar.xz"
	clangformat_cmd_src := exec.Command("curl", "-L", clangformat_src_url, "-o", "source.tar.gz")
	err = clangformat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clangformat_cmd_direct := exec.Command("./binary")
	err = clangformat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
}
