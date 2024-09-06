package main

// ClangFormatAT11 - Formatting tools for C, C++, Obj-C, Java, JavaScript, TypeScript
// Homepage: https://clang.llvm.org/docs/ClangFormat.html

import (
	"fmt"
	
	"os/exec"
)

func installClangFormatAT11() {
	// Método 1: Descargar y extraer .tar.gz
	clangformatat11_tar_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-11.1.0/llvm-11.1.0.src.tar.xz"
	clangformatat11_cmd_tar := exec.Command("curl", "-L", clangformatat11_tar_url, "-o", "package.tar.gz")
	err := clangformatat11_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clangformatat11_zip_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-11.1.0/llvm-11.1.0.src.tar.xz"
	clangformatat11_cmd_zip := exec.Command("curl", "-L", clangformatat11_zip_url, "-o", "package.zip")
	err = clangformatat11_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clangformatat11_bin_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-11.1.0/llvm-11.1.0.src.tar.xz"
	clangformatat11_cmd_bin := exec.Command("curl", "-L", clangformatat11_bin_url, "-o", "binary.bin")
	err = clangformatat11_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clangformatat11_src_url := "https://github.com/llvm/llvm-project/releases/download/llvmorg-11.1.0/llvm-11.1.0.src.tar.xz"
	clangformatat11_cmd_src := exec.Command("curl", "-L", clangformatat11_src_url, "-o", "source.tar.gz")
	err = clangformatat11_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clangformatat11_cmd_direct := exec.Command("./binary")
	err = clangformatat11_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
