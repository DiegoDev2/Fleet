package main

// ClangUml - Customizable automatic UML diagram generator for C++ based on Clang
// Homepage: https://github.com/bkryza/clang-uml

import (
	"fmt"
	
	"os/exec"
)

func installClangUml() {
	// Método 1: Descargar y extraer .tar.gz
	clanguml_tar_url := "https://github.com/bkryza/clang-uml/archive/refs/tags/0.5.4.tar.gz"
	clanguml_cmd_tar := exec.Command("curl", "-L", clanguml_tar_url, "-o", "package.tar.gz")
	err := clanguml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clanguml_zip_url := "https://github.com/bkryza/clang-uml/archive/refs/tags/0.5.4.zip"
	clanguml_cmd_zip := exec.Command("curl", "-L", clanguml_zip_url, "-o", "package.zip")
	err = clanguml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clanguml_bin_url := "https://github.com/bkryza/clang-uml/archive/refs/tags/0.5.4.bin"
	clanguml_cmd_bin := exec.Command("curl", "-L", clanguml_bin_url, "-o", "binary.bin")
	err = clanguml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clanguml_src_url := "https://github.com/bkryza/clang-uml/archive/refs/tags/0.5.4.src.tar.gz"
	clanguml_cmd_src := exec.Command("curl", "-L", clanguml_src_url, "-o", "source.tar.gz")
	err = clanguml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clanguml_cmd_direct := exec.Command("./binary")
	err = clanguml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: yaml-cpp")
	exec.Command("latte", "install", "yaml-cpp").Run()
}
