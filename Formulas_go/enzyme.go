package main

// Enzyme - High-performance automatic differentiation of LLVM
// Homepage: https://enzyme.mit.edu

import (
	"fmt"
	
	"os/exec"
)

func installEnzyme() {
	// Método 1: Descargar y extraer .tar.gz
	enzyme_tar_url := "https://github.com/EnzymeAD/Enzyme/archive/refs/tags/v0.0.148.tar.gz"
	enzyme_cmd_tar := exec.Command("curl", "-L", enzyme_tar_url, "-o", "package.tar.gz")
	err := enzyme_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	enzyme_zip_url := "https://github.com/EnzymeAD/Enzyme/archive/refs/tags/v0.0.148.zip"
	enzyme_cmd_zip := exec.Command("curl", "-L", enzyme_zip_url, "-o", "package.zip")
	err = enzyme_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	enzyme_bin_url := "https://github.com/EnzymeAD/Enzyme/archive/refs/tags/v0.0.148.bin"
	enzyme_cmd_bin := exec.Command("curl", "-L", enzyme_bin_url, "-o", "binary.bin")
	err = enzyme_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	enzyme_src_url := "https://github.com/EnzymeAD/Enzyme/archive/refs/tags/v0.0.148.src.tar.gz"
	enzyme_cmd_src := exec.Command("curl", "-L", enzyme_src_url, "-o", "source.tar.gz")
	err = enzyme_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	enzyme_cmd_direct := exec.Command("./binary")
	err = enzyme_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: llvm")
exec.Command("latte", "install", "llvm")
}
