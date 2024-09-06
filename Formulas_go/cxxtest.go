package main

// Cxxtest - C++ unit testing framework similar to JUnit, CppUnit and xUnit
// Homepage: https://github.com/CxxTest/cxxtest

import (
	"fmt"
	
	"os/exec"
)

func installCxxtest() {
	// Método 1: Descargar y extraer .tar.gz
	cxxtest_tar_url := "https://github.com/CxxTest/cxxtest/releases/download/4.4/cxxtest-4.4.tar.gz"
	cxxtest_cmd_tar := exec.Command("curl", "-L", cxxtest_tar_url, "-o", "package.tar.gz")
	err := cxxtest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cxxtest_zip_url := "https://github.com/CxxTest/cxxtest/releases/download/4.4/cxxtest-4.4.zip"
	cxxtest_cmd_zip := exec.Command("curl", "-L", cxxtest_zip_url, "-o", "package.zip")
	err = cxxtest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cxxtest_bin_url := "https://github.com/CxxTest/cxxtest/releases/download/4.4/cxxtest-4.4.bin"
	cxxtest_cmd_bin := exec.Command("curl", "-L", cxxtest_bin_url, "-o", "binary.bin")
	err = cxxtest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cxxtest_src_url := "https://github.com/CxxTest/cxxtest/releases/download/4.4/cxxtest-4.4.src.tar.gz"
	cxxtest_cmd_src := exec.Command("curl", "-L", cxxtest_src_url, "-o", "source.tar.gz")
	err = cxxtest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cxxtest_cmd_direct := exec.Command("./binary")
	err = cxxtest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
