package main

// Cppcheck - Static analysis of C and C++ code
// Homepage: https://sourceforge.net/projects/cppcheck/

import (
	"fmt"
	
	"os/exec"
)

func installCppcheck() {
	// Método 1: Descargar y extraer .tar.gz
	cppcheck_tar_url := "https://github.com/danmar/cppcheck/archive/refs/tags/2.15.0.tar.gz"
	cppcheck_cmd_tar := exec.Command("curl", "-L", cppcheck_tar_url, "-o", "package.tar.gz")
	err := cppcheck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cppcheck_zip_url := "https://github.com/danmar/cppcheck/archive/refs/tags/2.15.0.zip"
	cppcheck_cmd_zip := exec.Command("curl", "-L", cppcheck_zip_url, "-o", "package.zip")
	err = cppcheck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cppcheck_bin_url := "https://github.com/danmar/cppcheck/archive/refs/tags/2.15.0.bin"
	cppcheck_cmd_bin := exec.Command("curl", "-L", cppcheck_bin_url, "-o", "binary.bin")
	err = cppcheck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cppcheck_src_url := "https://github.com/danmar/cppcheck/archive/refs/tags/2.15.0.src.tar.gz"
	cppcheck_cmd_src := exec.Command("curl", "-L", cppcheck_src_url, "-o", "source.tar.gz")
	err = cppcheck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cppcheck_cmd_direct := exec.Command("./binary")
	err = cppcheck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pcre")
	exec.Command("latte", "install", "pcre").Run()
	fmt.Println("Instalando dependencia: tinyxml2")
	exec.Command("latte", "install", "tinyxml2").Run()
}
