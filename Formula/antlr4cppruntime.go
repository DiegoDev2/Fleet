package main

// Antlr4CppRuntime - ANother Tool for Language Recognition C++ Runtime Library
// Homepage: https://www.antlr.org/

import (
	"fmt"
	
	"os/exec"
)

func installAntlr4CppRuntime() {
	// Método 1: Descargar y extraer .tar.gz
	antlr4cppruntime_tar_url := "https://www.antlr.org/download/antlr4-cpp-runtime-4.13.2-source.zip"
	antlr4cppruntime_cmd_tar := exec.Command("curl", "-L", antlr4cppruntime_tar_url, "-o", "package.tar.gz")
	err := antlr4cppruntime_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	antlr4cppruntime_zip_url := "https://www.antlr.org/download/antlr4-cpp-runtime-4.13.2-source.zip"
	antlr4cppruntime_cmd_zip := exec.Command("curl", "-L", antlr4cppruntime_zip_url, "-o", "package.zip")
	err = antlr4cppruntime_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	antlr4cppruntime_bin_url := "https://www.antlr.org/download/antlr4-cpp-runtime-4.13.2-source.zip"
	antlr4cppruntime_cmd_bin := exec.Command("curl", "-L", antlr4cppruntime_bin_url, "-o", "binary.bin")
	err = antlr4cppruntime_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	antlr4cppruntime_src_url := "https://www.antlr.org/download/antlr4-cpp-runtime-4.13.2-source.zip"
	antlr4cppruntime_cmd_src := exec.Command("curl", "-L", antlr4cppruntime_src_url, "-o", "source.tar.gz")
	err = antlr4cppruntime_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	antlr4cppruntime_cmd_direct := exec.Command("./binary")
	err = antlr4cppruntime_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
