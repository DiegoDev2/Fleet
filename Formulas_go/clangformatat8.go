package main

// ClangFormatAT8 - Formatting tools for C, C++, Obj-C, Java, JavaScript, TypeScript
// Homepage: https://clang.llvm.org/docs/ClangFormat.html

import (
	"fmt"
	
	"os/exec"
)

func installClangFormatAT8() {
	// Método 1: Descargar y extraer .tar.gz
	clangformatat8_tar_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-8.0.1/llvm-8.0.1.src.tar.xz"
	clangformatat8_cmd_tar := exec.Command("curl", "-L", clangformatat8_tar_url, "-o", "package.tar.gz")
	err := clangformatat8_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clangformatat8_zip_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-8.0.1/llvm-8.0.1.src.tar.xz"
	clangformatat8_cmd_zip := exec.Command("curl", "-L", clangformatat8_zip_url, "-o", "package.zip")
	err = clangformatat8_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clangformatat8_bin_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-8.0.1/llvm-8.0.1.src.tar.xz"
	clangformatat8_cmd_bin := exec.Command("curl", "-L", clangformatat8_bin_url, "-o", "binary.bin")
	err = clangformatat8_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clangformatat8_src_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-8.0.1/llvm-8.0.1.src.tar.xz"
	clangformatat8_cmd_src := exec.Command("curl", "-L", clangformatat8_src_url, "-o", "source.tar.gz")
	err = clangformatat8_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clangformatat8_cmd_direct := exec.Command("./binary")
	err = clangformatat8_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
