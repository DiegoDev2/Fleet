package main

// C3c - Compiler for the C3 language
// Homepage: https://github.com/c3lang/c3c

import (
	"fmt"
	
	"os/exec"
)

func installC3c() {
	// Método 1: Descargar y extraer .tar.gz
	c3c_tar_url := "https://github.com/c3lang/c3c/archive/refs/tags/v0.6.1.tar.gz"
	c3c_cmd_tar := exec.Command("curl", "-L", c3c_tar_url, "-o", "package.tar.gz")
	err := c3c_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	c3c_zip_url := "https://github.com/c3lang/c3c/archive/refs/tags/v0.6.1.zip"
	c3c_cmd_zip := exec.Command("curl", "-L", c3c_zip_url, "-o", "package.zip")
	err = c3c_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	c3c_bin_url := "https://github.com/c3lang/c3c/archive/refs/tags/v0.6.1.bin"
	c3c_cmd_bin := exec.Command("curl", "-L", c3c_bin_url, "-o", "binary.bin")
	err = c3c_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	c3c_src_url := "https://github.com/c3lang/c3c/archive/refs/tags/v0.6.1.src.tar.gz"
	c3c_cmd_src := exec.Command("curl", "-L", c3c_src_url, "-o", "source.tar.gz")
	err = c3c_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	c3c_cmd_direct := exec.Command("./binary")
	err = c3c_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: z3")
	exec.Command("latte", "install", "z3").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
