package main

// Wllvm - Toolkit for building whole-program LLVM bitcode files
// Homepage: https://pypi.org/project/wllvm/

import (
	"fmt"
	
	"os/exec"
)

func installWllvm() {
	// Método 1: Descargar y extraer .tar.gz
	wllvm_tar_url := "https://files.pythonhosted.org/packages/4b/df/31d7519052bc21d0e9771e9a6540d6310bfb13bae7dacde060d8f647b8d3/wllvm-1.3.1.tar.gz"
	wllvm_cmd_tar := exec.Command("curl", "-L", wllvm_tar_url, "-o", "package.tar.gz")
	err := wllvm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wllvm_zip_url := "https://files.pythonhosted.org/packages/4b/df/31d7519052bc21d0e9771e9a6540d6310bfb13bae7dacde060d8f647b8d3/wllvm-1.3.1.zip"
	wllvm_cmd_zip := exec.Command("curl", "-L", wllvm_zip_url, "-o", "package.zip")
	err = wllvm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wllvm_bin_url := "https://files.pythonhosted.org/packages/4b/df/31d7519052bc21d0e9771e9a6540d6310bfb13bae7dacde060d8f647b8d3/wllvm-1.3.1.bin"
	wllvm_cmd_bin := exec.Command("curl", "-L", wllvm_bin_url, "-o", "binary.bin")
	err = wllvm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wllvm_src_url := "https://files.pythonhosted.org/packages/4b/df/31d7519052bc21d0e9771e9a6540d6310bfb13bae7dacde060d8f647b8d3/wllvm-1.3.1.src.tar.gz"
	wllvm_cmd_src := exec.Command("curl", "-L", wllvm_src_url, "-o", "source.tar.gz")
	err = wllvm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wllvm_cmd_direct := exec.Command("./binary")
	err = wllvm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
