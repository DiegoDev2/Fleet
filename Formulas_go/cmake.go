package main

// Cmake - Cross-platform make
// Homepage: https://www.cmake.org/

import (
	"fmt"
	
	"os/exec"
)

func installCmake() {
	// Método 1: Descargar y extraer .tar.gz
	cmake_tar_url := "https://github.com/Kitware/CMake/releases/download/v3.30.3/cmake-3.30.3.tar.gz"
	cmake_cmd_tar := exec.Command("curl", "-L", cmake_tar_url, "-o", "package.tar.gz")
	err := cmake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmake_zip_url := "https://github.com/Kitware/CMake/releases/download/v3.30.3/cmake-3.30.3.zip"
	cmake_cmd_zip := exec.Command("curl", "-L", cmake_zip_url, "-o", "package.zip")
	err = cmake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmake_bin_url := "https://github.com/Kitware/CMake/releases/download/v3.30.3/cmake-3.30.3.bin"
	cmake_cmd_bin := exec.Command("curl", "-L", cmake_bin_url, "-o", "binary.bin")
	err = cmake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmake_src_url := "https://github.com/Kitware/CMake/releases/download/v3.30.3/cmake-3.30.3.src.tar.gz"
	cmake_cmd_src := exec.Command("curl", "-L", cmake_src_url, "-o", "source.tar.gz")
	err = cmake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmake_cmd_direct := exec.Command("./binary")
	err = cmake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}
