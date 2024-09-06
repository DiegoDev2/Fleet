package main

// UnittestCpp - Unit testing framework for C++
// Homepage: https://github.com/unittest-cpp/unittest-cpp

import (
	"fmt"
	
	"os/exec"
)

func installUnittestCpp() {
	// Método 1: Descargar y extraer .tar.gz
	unittestcpp_tar_url := "https://github.com/unittest-cpp/unittest-cpp/releases/download/v2.0.0/unittest-cpp-2.0.0.tar.gz"
	unittestcpp_cmd_tar := exec.Command("curl", "-L", unittestcpp_tar_url, "-o", "package.tar.gz")
	err := unittestcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	unittestcpp_zip_url := "https://github.com/unittest-cpp/unittest-cpp/releases/download/v2.0.0/unittest-cpp-2.0.0.zip"
	unittestcpp_cmd_zip := exec.Command("curl", "-L", unittestcpp_zip_url, "-o", "package.zip")
	err = unittestcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	unittestcpp_bin_url := "https://github.com/unittest-cpp/unittest-cpp/releases/download/v2.0.0/unittest-cpp-2.0.0.bin"
	unittestcpp_cmd_bin := exec.Command("curl", "-L", unittestcpp_bin_url, "-o", "binary.bin")
	err = unittestcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	unittestcpp_src_url := "https://github.com/unittest-cpp/unittest-cpp/releases/download/v2.0.0/unittest-cpp-2.0.0.src.tar.gz"
	unittestcpp_cmd_src := exec.Command("curl", "-L", unittestcpp_src_url, "-o", "source.tar.gz")
	err = unittestcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	unittestcpp_cmd_direct := exec.Command("./binary")
	err = unittestcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
