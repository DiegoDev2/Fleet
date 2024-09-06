package main

// Cppinsights - See your source code with the eyes of a compiler
// Homepage: https://cppinsights.io/

import (
	"fmt"
	
	"os/exec"
)

func installCppinsights() {
	// Método 1: Descargar y extraer .tar.gz
	cppinsights_tar_url := "https://github.com/andreasfertig/cppinsights/archive/refs/tags/v_17.0.tar.gz"
	cppinsights_cmd_tar := exec.Command("curl", "-L", cppinsights_tar_url, "-o", "package.tar.gz")
	err := cppinsights_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cppinsights_zip_url := "https://github.com/andreasfertig/cppinsights/archive/refs/tags/v_17.0.zip"
	cppinsights_cmd_zip := exec.Command("curl", "-L", cppinsights_zip_url, "-o", "package.zip")
	err = cppinsights_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cppinsights_bin_url := "https://github.com/andreasfertig/cppinsights/archive/refs/tags/v_17.0.bin"
	cppinsights_cmd_bin := exec.Command("curl", "-L", cppinsights_bin_url, "-o", "binary.bin")
	err = cppinsights_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cppinsights_src_url := "https://github.com/andreasfertig/cppinsights/archive/refs/tags/v_17.0.src.tar.gz"
	cppinsights_cmd_src := exec.Command("curl", "-L", cppinsights_src_url, "-o", "source.tar.gz")
	err = cppinsights_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cppinsights_cmd_direct := exec.Command("./binary")
	err = cppinsights_cmd_direct.Run()
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
