package main

// Clingo - ASP system to ground and solve logic programs
// Homepage: https://potassco.org/clingo/

import (
	"fmt"
	
	"os/exec"
)

func installClingo() {
	// Método 1: Descargar y extraer .tar.gz
	clingo_tar_url := "https://github.com/potassco/clingo/archive/refs/tags/v5.7.1.tar.gz"
	clingo_cmd_tar := exec.Command("curl", "-L", clingo_tar_url, "-o", "package.tar.gz")
	err := clingo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clingo_zip_url := "https://github.com/potassco/clingo/archive/refs/tags/v5.7.1.zip"
	clingo_cmd_zip := exec.Command("curl", "-L", clingo_zip_url, "-o", "package.zip")
	err = clingo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clingo_bin_url := "https://github.com/potassco/clingo/archive/refs/tags/v5.7.1.bin"
	clingo_cmd_bin := exec.Command("curl", "-L", clingo_bin_url, "-o", "binary.bin")
	err = clingo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clingo_src_url := "https://github.com/potassco/clingo/archive/refs/tags/v5.7.1.src.tar.gz"
	clingo_cmd_src := exec.Command("curl", "-L", clingo_src_url, "-o", "source.tar.gz")
	err = clingo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clingo_cmd_direct := exec.Command("./binary")
	err = clingo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: re2c")
exec.Command("latte", "install", "re2c")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: cffi")
exec.Command("latte", "install", "cffi")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
