package main

// Numcpp - C++ implementation of the Python Numpy library
// Homepage: https://dpilger26.github.io/NumCpp

import (
	"fmt"
	
	"os/exec"
)

func installNumcpp() {
	// Método 1: Descargar y extraer .tar.gz
	numcpp_tar_url := "https://github.com/dpilger26/NumCpp/archive/refs/tags/Version_2.12.1.tar.gz"
	numcpp_cmd_tar := exec.Command("curl", "-L", numcpp_tar_url, "-o", "package.tar.gz")
	err := numcpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	numcpp_zip_url := "https://github.com/dpilger26/NumCpp/archive/refs/tags/Version_2.12.1.zip"
	numcpp_cmd_zip := exec.Command("curl", "-L", numcpp_zip_url, "-o", "package.zip")
	err = numcpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	numcpp_bin_url := "https://github.com/dpilger26/NumCpp/archive/refs/tags/Version_2.12.1.bin"
	numcpp_cmd_bin := exec.Command("curl", "-L", numcpp_bin_url, "-o", "binary.bin")
	err = numcpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	numcpp_src_url := "https://github.com/dpilger26/NumCpp/archive/refs/tags/Version_2.12.1.src.tar.gz"
	numcpp_cmd_src := exec.Command("curl", "-L", numcpp_src_url, "-o", "source.tar.gz")
	err = numcpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	numcpp_cmd_direct := exec.Command("./binary")
	err = numcpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
}
