package main

// Clangql - Run a SQL like language to perform queries on C/C++ files
// Homepage: https://github.com/AmrDeveloper/ClangQL

import (
	"fmt"
	
	"os/exec"
)

func installClangql() {
	// Método 1: Descargar y extraer .tar.gz
	clangql_tar_url := "https://github.com/AmrDeveloper/ClangQL/archive/refs/tags/0.6.0.tar.gz"
	clangql_cmd_tar := exec.Command("curl", "-L", clangql_tar_url, "-o", "package.tar.gz")
	err := clangql_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clangql_zip_url := "https://github.com/AmrDeveloper/ClangQL/archive/refs/tags/0.6.0.zip"
	clangql_cmd_zip := exec.Command("curl", "-L", clangql_zip_url, "-o", "package.zip")
	err = clangql_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clangql_bin_url := "https://github.com/AmrDeveloper/ClangQL/archive/refs/tags/0.6.0.bin"
	clangql_cmd_bin := exec.Command("curl", "-L", clangql_bin_url, "-o", "binary.bin")
	err = clangql_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clangql_src_url := "https://github.com/AmrDeveloper/ClangQL/archive/refs/tags/0.6.0.src.tar.gz"
	clangql_cmd_src := exec.Command("curl", "-L", clangql_src_url, "-o", "source.tar.gz")
	err = clangql_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clangql_cmd_direct := exec.Command("./binary")
	err = clangql_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
}
